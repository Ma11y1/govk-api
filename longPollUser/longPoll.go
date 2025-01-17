package longPollUser

import (
	"context"
	"fmt"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	internalErrors "go-vk-sdk/errors"
	"go-vk-sdk/logger"
	"go-vk-sdk/request"
	"sync/atomic"
)

// Doc: https://dev.vk.com/ru/api/user-long-poll/getting-started

type Server struct {
	URL string `json:"url"`
	Key string `json:"key"`
	Ts  int    `json:"ts"`
}

type LongPoll struct {
	api        *api.API
	user       actor.Actor
	mode       ExtraOptionsMode
	version    int
	url        string
	key        string
	ts         int
	wait       int
	chanUpdate chan *EventUpdate
	isRunning  int32

	req             *request.LongPollUserRequest              // cache
	reqUpdateServer *request.MessagesGetLongPollServerRequest // cache
}

func NewLongPoll(a *api.API, user actor.Actor, mode ExtraOptionsMode) *LongPoll {
	return &LongPoll{
		api:             a,
		user:            user,
		mode:            mode,
		version:         3,
		url:             "",
		key:             "",
		ts:              -1,
		wait:            25,
		chanUpdate:      make(chan *EventUpdate, 2),
		isRunning:       0,
		req:             request.NewLongPollUserRequest(a, "").Wait(90).Mode(int(mode)).Version(3),
		reqUpdateServer: request.NewMessagesGetLongPollServerRequest(a, user).LpVersion(3),
	}
}

func NewLongPollServer(a *api.API, user actor.Actor, mode ExtraOptionsMode, server *Server) *LongPoll {
	lp := NewLongPoll(a, user, mode)
	lp.url = server.URL
	lp.key = server.Key
	lp.ts = server.Ts
	return lp
}

func (l *LongPoll) UpdateServer(isUpdateTs bool) error {
	serverSettings, err := l.reqUpdateServer.Exec(context.Background())
	if err != nil {
		return err
	}

	if serverSettings.Response.Key == "" {
		return internalErrors.ErrorLog("LongPollUser.LongPoll.UpdateServer()", "Response server settings is empty")
	}

	l.key = serverSettings.Response.Key
	l.url = serverSettings.Response.Server

	if isUpdateTs {
		l.ts = serverSettings.Response.Ts
	}

	l.req.
		Key(l.key).
		Ts(l.ts).
		SetURL(l.url)

	return nil
}

func (l *LongPoll) Run(ctx context.Context) error {
	if l.url == "" || l.key == "" {
		return internalErrors.ErrorLog("LongPollUser.LongPoll.Run()", "Server is undefined")
	}

	if !atomic.CompareAndSwapInt32(&l.isRunning, 0, 1) {
		return internalErrors.ErrorLog("LongPollUser.LongPoll.Run()", "Long poll user is already running")
	}

	defer atomic.StoreInt32(&l.isRunning, 0)

	logger.Log("LongPollGroup.LongPoll.Run()", "Long poll user server is running at url "+l.url)

	for atomic.LoadInt32(&l.isRunning) == 1 {
		select {
		case _, ok := <-ctx.Done():
			if !ok {
				return ctx.Err()
			}
		default:
			resp, err := l.req.Exec(context.Background())
			if err != nil {
				logger.Log("LongPollUser.LongPoll.Run()", err.Error())
				continue
			}

			switch FailedType(resp.Failed) {
			case 0, FailedTypeOutdatedStory:
				l.ts = resp.Ts
			case FailedTypeExpiredKey:
				err = l.UpdateServer(false)
			case FailedTypeOutdatedUserInfo:
				err = l.UpdateServer(true)
			case FailedTypeInvalidVersion:
				err = NotValidVersionError
			default:
				err = &FailedError{Code: resp.Failed}
			}
			if err != nil {
				logger.Log("LongPollUser.LongPoll.Run()", err.Error())
				continue
			}

			for _, data := range resp.Updates {
				event, err := newEventUpdate(data, l.mode)
				if err != nil {
					logger.Log("LongPollUser.LongPoll.Run()", err.Error())
					continue
				}

				l.chanUpdate <- event
			}
		}
	}

	logger.Log("LongPollGroup.LongPoll.Run()", "Long poll user server is stopped at url "+l.url)

	return nil
}

func (l *LongPoll) Stop() error {
	atomic.StoreInt32(&l.isRunning, 0)
	return nil
}

func (l *LongPoll) Updates() chan *EventUpdate {
	return l.chanUpdate
}

func (l *LongPoll) SetServer(server *Server) error {
	if server == nil || server.URL == "" || server.Key == "" {
		return internalErrors.ErrorLog("LongPollUser.LongPoll.SetServer()", fmt.Sprintf("Invalid server configuration %+v", server))
	}

	l.url = server.URL
	l.key = server.Key
	l.ts = server.Ts
	l.req.
		Key(l.key).
		Ts(l.ts).
		SetURL(l.url)

	return nil
}

func (l *LongPoll) SetMode(mode ExtraOptionsMode) {
	l.mode = mode
	l.req.Mode(int(mode))
}

// SetWait wait > 0 and < 90
func (l *LongPoll) SetWait(wait int) error {
	if wait <= 0 || wait > 90 {
		return internalErrors.ErrorLog("LongPollUser.LongPoll.SetWait()", fmt.Sprintf("Invalid wait time: %d, must be between 1 and 90", wait))
	}

	l.wait = wait
	l.req.Wait(wait)

	return nil
}

// SetVersion actually 3
func (l *LongPoll) SetVersion(v int) {
	if v <= 0 || v > 3 {
		return
	}
	l.version = v
	l.req.Version(v)
}

func (l *LongPoll) IsRunning() bool {
	return atomic.LoadInt32(&l.isRunning) == 1
}

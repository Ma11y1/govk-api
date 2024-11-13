package callback

import (
	"context"
	"encoding/json"
	"fmt"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	internalErrors "go-vk-sdk/errors"
	"go-vk-sdk/events"
	"go-vk-sdk/logger"
	"go-vk-sdk/objects"
	"go-vk-sdk/request"
	"go-vk-sdk/transport"
	"net/http"
	"strconv"
)

// Doc: https://dev.vk.com/ru/api/callback/getting-started

type ServerStatus string

const (
	ServerStatusOk           ServerStatus = "ok"
	ServerStatusWait         ServerStatus = "wait"
	ServerStatusFailed       ServerStatus = "failed"
	ServerStatusUnconfigured ServerStatus = "unconfigured"
)

type Callback struct {
	server           transport.CallbackServer
	api              *api.API
	actor            actor.Actor
	eventEmitter     *events.EventEmitter[events.EventType, *events.EventCallback]
	ConfirmationKey  string
	confirmationKeys map[int]string
	SecretKey        string
	secretKeys       map[int]string
}

func NewCallback(api *api.API, actor actor.Actor, address, handlePath string) *Callback {
	if address == "" || handlePath == "" {
		panic(internalErrors.ErrorLog("Callback.NewCallback()", "invalid value address or handlePath"))
	}

	m := http.NewServeMux()

	callback := &Callback{
		api:              api,
		actor:            actor,
		server:           transport.NewBaseCallbackServerWithHandler(address, m),
		eventEmitter:     events.NewEventEmitter[events.EventType, *events.EventCallback](),
		confirmationKeys: make(map[int]string),
		secretKeys:       make(map[int]string),
	}

	m.HandleFunc(handlePath, callback.handle)

	return callback
}

func NewCallbackServer(api *api.API, actor actor.Actor, server transport.CallbackServer) *Callback {
	return &Callback{
		api:              api,
		actor:            actor,
		server:           server,
		eventEmitter:     events.NewEventEmitter[events.EventType, *events.EventCallback](),
		confirmationKeys: make(map[int]string),
		secretKeys:       make(map[int]string),
	}
}

func (c *Callback) Run() error {
	err := c.server.Run()
	if err != nil {
		return internalErrors.ErrorLog("Callback.Run()", "Failed to start callback server: "+err.Error())
	}
	return nil
}

func (c *Callback) Stop(ctx context.Context) error {
	err := c.server.Stop(ctx)
	if err != nil {
		return internalErrors.ErrorLog("Callback.Stop()", "Failed to stop callback server: "+err.Error())
	}
	return nil
}

// handle base handler at a given path
func (c *Callback) handle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var updateEvent events.EventUpdate
	err := decoder.Decode(&updateEvent)
	if err != nil {
		logger.Log("Callback.handle()", "error JSON decode result body: "+err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	secretKey, ok := c.secretKeys[updateEvent.GroupID]
	if !ok {
		secretKey = c.SecretKey
	}

	if secretKey != "" && updateEvent.Secret != secretKey {
		logger.Log("Callback.handle()", fmt.Sprintf("Bad secret key %s for grouod id %d", secretKey, updateEvent.GroupID))
		http.Error(w, "Bad secret", http.StatusForbidden)

		return
	}

	if updateEvent.Type == events.EventTypeConfirmation {
		key, exists := c.confirmationKeys[updateEvent.GroupID]
		if exists && key != "" {
			_, err = w.Write([]byte(key))
			if err != nil {
				logger.Log("Callback.handle()", "failed to write confirmation key to response writer: "+err.Error())
			}
		} else {
			c.eventEmitter.Emit(events.EventTypeConfirmation, &events.EventCallback{Event: &events.EventConfirmation{
				Response: w,
				GroupID:  updateEvent.GroupID,
			}})
		}
		return
	}

	event, err := events.NewEvent(&updateEvent)
	if err != nil {
		logger.Log("Callback.handle()", "failed to create event: "+err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	callbackEvent := &events.EventCallback{Event: event}

	callbackEvent.RetryCounter, _ = strconv.Atoi(r.Header.Get("X-Retry-Counter"))

	c.eventEmitter.Emit(updateEvent.Type, callbackEvent)

	if callbackEvent.Error != nil {
		logger.Log("Callback.handle()", "failed to handle callback event: "+callbackEvent.Error.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if callbackEvent.IsRemove {
		_, _ = w.Write([]byte("remove"))
		logger.Log("Callback.handle()", fmt.Sprintf("group %d server was deleted", updateEvent.GroupID))
		return
	}

	if callbackEvent.Code != 0 {
		w.Header().Set("Retry-After", callbackEvent.Date.Format(http.TimeFormat))
		http.Error(w, http.StatusText(callbackEvent.Code), callbackEvent.Code)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (c *Callback) AddEventListener(event events.EventType, listener *events.EventListener[*events.EventCallback]) {
	if event == "" || listener == nil {
		logger.Log("Callback.AddEventListener()", "attempted to add nil event or listener")
		return
	}
	c.eventEmitter.On(event, listener)
}

func (c *Callback) RemoveEventListener(event events.EventType, listener *events.EventListener[*events.EventCallback]) {
	if event == "" || listener == nil {
		logger.Log("Callback.RemoveEventListener()", "attempted to remove nil event or listener")
		return
	}
	c.eventEmitter.Off(event, listener)
}

func (c *Callback) ClearEventListeners(event events.EventType) {
	c.eventEmitter.Clear(event)
}

func (c *Callback) SetDefaultHandler(path string) error {
	return c.SetHandleFunc(path, c.handle)
}

func (c *Callback) SetHandler(path string, handler http.Handler) error {
	if path == "" {
		return internalErrors.ErrorLog("Callback.SetHandler()", "invalid value path")
	}

	if path[0] != '/' {
		return internalErrors.ErrorLog("Callback.SetHandler()", "The first character of the path must be '/': "+path)
	}

	m := http.NewServeMux()
	m.Handle(path, handler)
	return c.server.SetHandler(m)
}

func (c *Callback) SetHandleFunc(path string, fn func(http.ResponseWriter, *http.Request)) error {
	if path == "" {
		return internalErrors.ErrorLog("Callback.SetHandleFunc()", "invalid value path")
	}

	if path[0] != '/' {
		return internalErrors.ErrorLog("Callback.SetHandleFunc()", "The first character of the path must be '/': "+path)
	}

	m := http.NewServeMux()
	m.HandleFunc(path, fn)
	return c.server.SetHandler(m)
}

// GetServers Retrieves information about servers for the Callback API in the community
func (c *Callback) GetServers(groupID int) ([]objects.GroupCallbackServer, error) {
	res, err := request.NewGroupsGetCallbackServersRequest(c.api, c.actor).
		GroupID(groupID).
		Exec(context.Background())
	if err != nil {
		return nil, internalErrors.ErrorLog("Callback.GetCallbackServers()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	return res.Response.Items, nil
}

// AddServer Adds a server for the Callback API to the community
//
//	After successful execution, returns the identifier of the added server in the server_id (integer) field
func (c *Callback) AddServer(groupID int, title, url, secret string) (int, error) {
	res, err := request.NewGroupsAddCallbackServerRequest(c.api, c.actor).
		GroupID(groupID).
		Title(title).
		URL(url).
		SecretKey(secret).
		Exec(context.Background())

	if err != nil {
		return -1, internalErrors.ErrorLog("Callback.AddServer()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	return res.Response.ServerID, nil
}

// DeleteServer Removes the server for the Callback API from the community
func (c *Callback) DeleteServer(serverID, groupID int) (bool, error) {
	res, err := request.NewGroupsDeleteCallbackServerRequest(c.api, c.actor).
		GroupID(groupID).
		ServerID(serverID).
		Exec(context.Background())

	if err != nil {
		return false, internalErrors.ErrorLog("Callback.DeleteServer()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	if res.Response == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

// SetSettings Allows you to set event notification settings in the Callback API
func (c *Callback) SetSettings(groupID, serverID int) (bool, error) {
	return c.SetSettingsEvents(groupID, serverID, c.eventEmitter.Keys())
}

// SetSettingsEvents Allows you to set event notification settings in the Callback API
func (c *Callback) SetSettingsEvents(groupID, serverID int, e []events.EventType) (bool, error) {
	request := request.NewGroupsSetCallbackSettingsRequest(c.api, c.actor).
		GroupID(groupID).
		ServerID(serverID).
		APIVersion(c.api.Version)

	for key, _ := range e {
		request.SetEvent(string(key), true)
	}

	res, err := request.Exec(context.Background())
	if err != nil {
		return false, internalErrors.ErrorLog("Callback.SetSettings()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	return true, nil
}

// GetConfirmationCode Allows you to get the string required to confirm the server address in the Callback API
func (c *Callback) GetConfirmationCode(groupID int) (string, error) {
	res, err := request.NewGroupsGetCallbackConfirmationCodeRequest(c.api, c.actor).
		GroupID(groupID).
		Exec(context.Background())

	if err != nil {
		return "", internalErrors.ErrorLog("Callback.GetConfirmationCode()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	return res.Response.Code, nil
}

func (c *Callback) IsRunning() bool {
	return c.server.IsRunning()
}
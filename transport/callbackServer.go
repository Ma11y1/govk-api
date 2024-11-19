package transport

import (
	"context"
	"errors"
	"fmt"
	internalErrors "go-vk-sdk/errors"
	"go-vk-sdk/logger"
	"net/http"
	"net/url"
)

type CallbackServer interface {
	Run() error
	Stop(context.Context) error
	GetURL() *url.URL
	SetURL(url.URL) error
	SetHandler(path string, handler http.Handler) error
	SetHandleFunc(path string, fn func(http.ResponseWriter, *http.Request)) error
	IsRunning() bool
}

type BaseCallbackServer struct {
	server    *http.Server
	handler   http.Handler
	url       *url.URL
	isRunning bool
}

func NewBaseCallbackServer(url *url.URL) *BaseCallbackServer {
	if url == nil || url.Host == "" {
		panic(internalErrors.ErrorLog("BaseCallbackServer.NewBaseCallbackServer()", "Invalid URL"))
	}
	return &BaseCallbackServer{
		server: &http.Server{
			Addr: url.Host,
		},
		handler:   nil,
		url:       url,
		isRunning: false,
	}
}

func (s *BaseCallbackServer) Run() error {
	if s.isRunning {
		return internalErrors.ErrorLog("BaseCallbackServer.Run()", "Server already running")
	}

	s.isRunning = true
	go func() {
		logger.Log("BaseCallbackServer.Run()", fmt.Sprintf("Server is running at url: %s", s.url.String()))

		if s.handler == nil {
			logger.Log("BaseCallbackServer.Run()", "Handler is undefined")
		}

		err := s.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log("BaseCallbackServer.Run()", "Error listen and serve server: "+err.Error())
			s.isRunning = false
			return
		}
	}()

	return nil
}

func (s *BaseCallbackServer) Stop(ctx context.Context) error {
	if !s.isRunning {
		return internalErrors.ErrorLog("BaseCallbackServer.Stop()", "Server is not running")
	}

	s.isRunning = false

	err := s.server.Shutdown(ctx)
	if err != nil {
		return internalErrors.ErrorLog("BaseCallbackServer.Stop()", "Server stop error: "+err.Error())
	}

	logger.Log("BaseCallbackServer.Stop()", "Server was stopped at url: "+s.url.String())

	s.server = &http.Server{
		Addr:    s.url.Host,
		Handler: s.handler,
	}

	return nil
}

func (s *BaseCallbackServer) GetURL() *url.URL {
	return s.url
}

func (s *BaseCallbackServer) SetURL(url url.URL) error {
	if s.isRunning {
		return internalErrors.ErrorLog("BaseCallbackServer.SetURL()", "Cannot change url while server is running")
	}

	s.url = &url
	s.server.Addr = url.Host

	return nil
}

func (s *BaseCallbackServer) SetHandler(path string, handler http.Handler) error {
	if s.isRunning {
		return internalErrors.ErrorLog("BaseCallbackServer.SetHandler()", "Cannot change handler while server is running")
	}

	if path == "" {
		return internalErrors.ErrorLog("BaseCallbackServer.SetHandler()", "invalid value path")
	}

	if path[0] != '/' {
		return internalErrors.ErrorLog("BaseCallbackServer.SetHandler()", "The first character of the path must be '/': "+path)
	}

	s.url.Path = path

	m := http.NewServeMux()
	m.Handle(path, handler)
	s.server.Handler = m
	s.handler = m

	return nil
}

func (s *BaseCallbackServer) SetHandleFunc(path string, fn func(http.ResponseWriter, *http.Request)) error {
	if s.isRunning {
		return internalErrors.ErrorLog("BaseCallbackServer.SetHandler()", "Cannot change handler while server is running")
	}

	if path == "" {
		return internalErrors.ErrorLog("BaseCallbackServer.SetHandleFunc()", "invalid value path")
	}

	if path[0] != '/' {
		return internalErrors.ErrorLog("BaseCallbackServer.SetHandleFunc()", "The first character of the path must be '/': "+path)
	}

	s.url.Path = path

	m := http.NewServeMux()
	m.HandleFunc(path, fn)
	s.server.Handler = m
	s.handler = m

	return nil
}

func (s *BaseCallbackServer) IsRunning() bool {
	return s.isRunning
}

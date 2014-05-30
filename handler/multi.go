package handler

import (
	"github.com/materials-commons/config"
)

// Keeps a list of all the handlers to use.
type multiHandler struct {
	handlers []config.Handler
}

// Multi takes a list of Handlers and returns a single Handler that calls them.
// The handlers are called in the order they are specified.
func Multi(handlers ...config.Handler) config.Handler {
	return &multiHandler{handlers: handlers}
}

// Init initializes each of the handlers. If any of the Handlers returns an error
// then Init returns an error. The results of calling Set or Get if Init returns
// an error are not specified.
func (h *multiHandler) Init() error {
	for _, handler := range h.handlers {
		if err := handler.Init(); err != nil {
			return err
		}
	}

	return nil
}

// Get iterates through each of the handlers in the order given in Multi. It stops
// when one of the handlers returns a value.
func (h *multiHandler) Get(key string) (interface{}, bool) {
	for _, handler := range h.handlers {
		if val, found := handler.Get(key); found {
			return val, true
		}
	}
	return nil, false
}

// Set iterates through each of the handlers in the order given in Multi. It stops
// when one of the handlers successfully sets the key value.
func (h *multiHandler) Set(key string, value interface{}) error {
	for _, handler := range h.handlers {
		if err := handler.Set(key, value); err == nil {
			return nil
		}
	}
	return config.ErrKeyNotSet
}

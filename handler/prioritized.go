package handler

import (
	"github.com/materials-commons/config"
)

type NamedHandler struct {
	Name    string
	Handler config.Handler
}

func NameHandler(name string, handler config.Handler) *NamedHandler {
	return &NamedHandler{
		Name:    name,
		Handler: handler,
	}
}

type prioritizedHandler struct {
	byName     map[string]config.Handler
	byPosition []*NamedHandler
}

func Prioritized(handlers ...*NamedHandler) config.Handler {
	phandler := &prioritizedHandler{
		byName:     make(map[string]config.Handler),
		byPosition: make([]*NamedHandler, len(handlers)),
	}

	for i, h := range handlers {
		phandler.byName[h.Name] = h.Handler
		phandler.byPosition[i] = h
	}

	return phandler
}

func (h *prioritizedHandler) Init() error {
	for _, handler := range h.byName {
		if err := handler.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (h *prioritizedHandler) Get(key string, args ...interface{}) (interface{}, error) {
	switch len(args) {
	case 0:
		return h.byPosition[0].Handler.Get(key)
	case 1:
		handler, err := h.getNamedHandler(args...)
		if err != nil {
			return nil, err
		}
		return handler.Get(key)
	default:
		return nil, config.ErrBadArgs
	}
}

func (h *prioritizedHandler) Set(key string, value interface{}, args ...interface{}) error {
	switch len(args) {
	case 0:
		return h.byPosition[0].Handler.Set(key, value)
	case 1:
		handler, err := h.getNamedHandler(args...)
		if err != nil {
			return err
		}
		return handler.Set(key, value)
	default:
		handler, err := h.getNamedHandler(args...)
		if err != nil {
			return err
		}
		args2 := args[1:]
		return handler.Set(key, value, args2...)
	}
}

func (h *prioritizedHandler) getNamedHandler(args ...interface{}) (config.Handler, error) {
	if handlerName, ok := args[0].(string); ok {
		handler, found := h.byName[handlerName]
		if !found {
			return nil, config.ErrBadArgs
		}
		return handler, nil
	}
	return nil, config.ErrBadArgs
}

func (h *prioritizedHandler) Args() bool {
	return true
}

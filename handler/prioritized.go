package handler

import (
	"github.com/materials-commons/config"
)

// A NamedHandler holds the name to use for a specific handler.
type NamedHandler struct {
	Name    string
	Handler config.Handler
}

// NameHandler is a convenience function for creating new instances
// of NamedHandler types.
func NameHandler(name string, handler config.Handler) *NamedHandler {
	return &NamedHandler{
		Name:    name,
		Handler: handler,
	}
}

type prioritizedHandler struct {
	byName     map[string]*NamedHandler // Look up handler by name
	byPosition []*NamedHandler          // Lookup handler by position
}

// Prioritized creates a new Prioritized Handler. A Prioritized Handler performs
// look ups in the order they were given. In addition the name of a handler can
// be passed as the last argument to a set or get. If this is done, then the
// named handler is used.
func Prioritized(handlers ...*NamedHandler) config.Handler {
	phandler := &prioritizedHandler{
		byName:     make(map[string]*NamedHandler),
		byPosition: make([]*NamedHandler, len(handlers)),
	}

	// Initialize the two ways of looking up a handler.
	for i, h := range handlers {
		phandler.byName[h.Name] = h
		phandler.byPosition[i] = h
	}

	return phandler
}

// Init initializes each of the handlers. If any handlers Init method returns
// an error then initialization stops and the error is returned.
func (h *prioritizedHandler) Init() error {
	for _, handler := range h.byName {
		if err := handler.Handler.Init(); err != nil {
			return err
		}
	}
	return nil
}

// Get looks up the given key. The first optional arg is the name of the handler to
// use. The optional args after the first are passed to the named handler. If no
// optional arg is passed then Get iterates through the handlers in the order they
// were given to the Prioritized constructor. It stops if one of these handlers
// returns nil for an error.
func (h *prioritizedHandler) Get(key string, args ...interface{}) (interface{}, error) {
	switch length := len(args); length {
	case 0:
		// No handler was given so go through all the handlers in
		// the order they were passed in.
		for _, nhandler := range h.byPosition {
			if val, err := nhandler.Handler.Get(key); err == nil {
				return val, nil
			}
		}
		return nil, config.ErrKeyNotFound
	default:
		// Lookup the handler.
		handler, err := h.getNamedHandler(args...)
		switch {
		case err != nil:
			return nil, err
		case length == 1:
			return handler.Get(key)
		default:
			// There was more than one arg, pass the other args on.
			otherArgs := args[1:]
			return handler.Get(key, otherArgs)
		}
	}
}

// Set set the value for the given key. The first optional arg is the name of the
// handler to use. The optional args after the first are passed to the named
// handler. If no optional arg is passed then Set iterates through the handlers
// in the order they were given to the Prioritized constructor. It stops if one
// of these handlers returns nil for an error.
func (h *prioritizedHandler) Set(key string, value interface{}, args ...interface{}) error {
	switch length := len(args); length {
	case 0:
		// No handler was given so go through all the handlers in
		// the order they were passed in.
		for _, nhandler := range h.byPosition {
			if err := nhandler.Handler.Set(key, value); err == nil {
				return nil
			}
		}
		return config.ErrKeyNotSet
	default:
		handler, err := h.getNamedHandler(args...)
		switch {
		case err != nil:
			return err
		case length == 1:
			return handler.Set(key, value)
		default:
			otherArgs := args[1:]
			return handler.Set(key, value, otherArgs)
		}
	}
}

// getNamedHandler takes the first arg, attempts to cast it to a string and looks up
// the named handler.
func (h *prioritizedHandler) getNamedHandler(args ...interface{}) (config.Handler, error) {
	if handlerName, ok := args[0].(string); ok {
		handler, found := h.byName[handlerName]
		if !found {
			return nil, config.ErrBadArgs
		}
		return handler.Handler, nil
	}
	return nil, config.ErrBadArgs
}

// Args returns true. Prioritized handlers always accept an optional argument that
// is the name of the handler to use.
func (h *prioritizedHandler) Args() bool {
	return true
}

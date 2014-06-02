package handler

import (
	"github.com/materials-commons/config/cfg"
)

type loaderHandler struct {
	handler cfg.Handler
	loader  cfg.Loader
}

// Loader returns a handler that reads the keys in from a loader.
func Loader(loader cfg.Loader) cfg.Handler {
	return &loaderHandler{
		handler: Map(),
		loader:  loader,
	}
}

// Init loads the keys by calling the loader.
func (h *loaderHandler) Init() error {
	m := h.handler.(*mapHandler)
	if err := h.loader.Load(&m.values); err != nil {
		return err
	}
	return nil
}

// Get retrieves keys loaded from the loader.
func (h *loaderHandler) Get(key string, args ...interface{}) (interface{}, error) {
	return h.handler.Get(key, args...)
}

// Set sets the value of keys. You can create new keys, or modify existing ones.
// Values are not persisted across runs.
func (h *loaderHandler) Set(key string, value interface{}, args ...interface{}) error {
	return h.handler.Set(key, value, args...)
}

// Args returns false. This handler doesn't accept additional arguments.
func (h *loaderHandler) Args() bool {
	return false
}

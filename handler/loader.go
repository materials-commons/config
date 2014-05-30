package handler

import (
	"github.com/materials-commons/config"
)

type loaderHandler struct {
	values map[string]interface{}
	loader config.Loader
}

// Loader returns a handler that reads the keys in from a loader.
func Loader(loader config.Loader) config.Handler {
	return &loaderHandler{
		loader: loader,
	}
}

// Init loads the keys by calling the loader.
func (h *loaderHandler) Init() error {
	var vals map[string]interface{}
	if err := h.loader.Load(&vals); err != nil {
		return err
	}
	h.values = vals
	return nil
}

// Get retrieves keys loaded from the loader.
func (h *loaderHandler) Get(key string) (interface{}, bool) {
	val, found := h.values[key]
	return val, found
}

// Set sets the value of keys. You can create new keys, or modify existing ones.
// Values are not persisted across runs.
func (h *loaderHandler) Set(key string, value interface{}) error {
	h.values[key] = value
	return nil
}

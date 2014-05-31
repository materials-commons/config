package handler

import (
	"github.com/materials-commons/config"
)

type mapHandler struct {
	values map[string]interface{}
}

func Map() config.Handler {
	return &mapHandler{values: make(map[string]interface{})}
}

func (h *mapHandler) Init() error {
	return nil
}

func (h *mapHandler) Get(key string, args ...interface{}) (interface{}, error) {
	if len(args) != 0 {
		return nil, config.ErrArgsNotSupported
	}
	val, found := h.values[key]
	if !found {
		return val, config.ErrKeyNotFound
	}
	return val, nil
}

// Set sets the value of keys. You can create new keys, or modify existing ones.
// Values are not persisted across runs.
func (h *mapHandler) Set(key string, value interface{}, args ...interface{}) error {
	if len(args) != 0 {
		return config.ErrArgsNotSupported
	}
	h.values[key] = value
	return nil
}

// Args returns false. This handler doesn't accept additional arguments.
func (h *mapHandler) Args() bool {
	return false
}

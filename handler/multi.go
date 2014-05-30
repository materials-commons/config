package handler

import (
	"github.com/materials-commons/config"
)

type multiHandler struct {
	handlers []config.Handler
}

func Multi(handlers... config.Handler) config.Handler {
	return &multiHandler{handlers: handlers}
}

func (h *multiHandler) Init() error {
	for _, handler := range h.handlers {
		handler.Init()
	}

	return nil
}

func (h *multiHandler) Get(key string) (interface{}, bool) {
	for _, handler := range h.handlers {
		if val, found := handler.Get(key); found {
			return val, true
		}
	}
	return nil, false
}

func (h *multiHandler) Set(key string, value interface{}) error {
	for _, handler := range h.handlers {
		if err := handler.Set(key, value); err == nil {
			return nil
		}
	}
	return config.ErrKeyNotSet
}















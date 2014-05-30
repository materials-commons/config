package config

import (
	"io"
	"os"
	"strings"
	"sync"
)

type Initer interface {
	Init() error
}

type Handler interface {
	Initer
	Getter
	Setter
}

type kvHandler struct {
	values map[string]interface{}
	loader Loader
}

func KVHandler(loader Loader) Handler {
	return &kvHandler{
		loader: loader,
	}
}

func (h *kvHandler) Init() error {
	var vals map[string]interface{}
	if err := h.loader.Load(&vals); err != nil {
		return err
	}
	h.values = vals
	return nil
}

func (h *kvHandler) Get(key string) (interface{}, bool) {
	val, found := h.values[key]
	return val, found
}

func (h *kvHandler) Set(key string, value interface{}) error {
	return nil
}

type envHandler struct{}

func EnvHandler() Handler {
	return &envHandler{}
}

func (h *envHandler) Init() error {
	return nil
}

func (h *envHandler) Get(key string) (interface{}, bool) {
	ukey := strings.ToUpper(key)
	val := os.Getenv(ukey)
	if val == "" {
		return val, false
	}
	return val, true
}

func (h *envHandler) Set(key string, value interface{}) error {
	ukey := strings.ToUpper(key)
	sval, err := toString(value)
	if err != nil {
		return ErrBadType
	}

	err = os.Setenv(ukey, sval)
	if err != nil {
		return err
	}

	return nil
}

type syncHandler struct {
	handler Handler
	loaded  bool
	mutex   sync.Mutex
}

func SyncHandler(handler Handler) Handler {
	return &syncHandler{handler: handler}
}

func (h *syncHandler) Init() error {
	defer h.mutex.Unlock()
	h.mutex.Lock()
	return h.handler.Init()
}

func (h *syncHandler) Get(key string) (interface{}, bool) {
	defer h.mutex.Unlock()
	h.mutex.Lock()
	return h.handler.Get(key)
}

func (h *syncHandler) Set(key string, value interface{}) error {
	defer h.mutex.Unlock()
	h.mutex.Lock()
	return h.handler.Set(key, value)
}


type multiHandler struct {
	handlers []Handler
}

func MultiHandler(handlers... Handler) Handler {
	return &multiHandler{handlers: handlers}
}

func (h *multiHandler) Init() error {
	for _, handler := h.handlers {
		handler.Init()
	}

	return nil
}

func (h *multiHandler) Get(key string) (interface{}, bool) {
	for _, handler := h.handlers {
		if val, found := handler.Get(key); found {
			return val, true
		}
	}
	return nil, false
}

func (h *multiHandler) Set(key string, value interface{}) error {
	for _, handler := h.handlers {
		if err := handler.Set(key, value); err == nil {
			return nil
		}
	}
	return ErrKeyNotSet
}

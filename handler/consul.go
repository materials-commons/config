package handler

import (
	"bytes"
	"encoding/gob"

	consul "github.com/armon/consul-api"
	"github.com/materials-commons/config/cfg"
)

type consulHandler struct {
	client *consul.Client
}

func Consul(client *consul.Client) cfg.Handler {
	return &consulHandler{
		client: client,
	}
}

func (h *consulHandler) Init() error {
	return nil
}

func (h *consulHandler) Get(key string, args ...interface{}) (interface{}, error) {
	if len(args) != 0 {
		return nil, cfg.ErrArgsNotSupported
	}

	kv, _, err := h.client.KV().Get(key, nil)
	if err != nil {
		return nil, cfg.ErrKeyNotFound
	}

	return kv.Value, nil
}

func (h *consulHandler) Set(key string, value interface{}, args ...interface{}) error {
	if len(args) != 0 {
		return cfg.ErrArgsNotSupported
	}

	asBytes, err := toBytes(value)
	if err != nil {
		return cfg.ErrBadType
	}

	kv := &consul.KVPair{
		Key:   key,
		Value: asBytes,
	}

	if _, err := h.client.KV().Put(kv, nil); err != nil {
		return cfg.ErrKeyNotSet
	}
	return nil
}

func (h *consulHandler) Args() bool {
	return false
}

func toBytes(value interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

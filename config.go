package config

import (
	"time"
)

type Config interface {
	Load(path string) error
	GetInt(key string) (int, error)
	GetString(key string) (string, error)
	GetTime(key string) (time.Time, error)
	GetBool(key string) (bool, error)
	Set(key string, value interface{}) error
	SetHandler(handler Handler)
}

type config struct {
	handler Handler
}

func New() Config {
	return nil
}

func (c *config) Load(path string) error {
	return c.handler.Load(path)
}

func (c *config) GetInt(key string) (int, error) {
	val, found := c.handler.Get(key)
	if found {
		return toInt(val)
	}

	return 0, ErrKeyNotFound
}

func (c *config) GetString(key string) (string, error) {
	val, found := c.handler.Get(key)
	if found {
		return toString(val)
	}

	return "", ErrKeyNotFound
}

func (c *config) GetTime(key string) (time.Time, error) {
	val, found := c.handler.Get(key)
	if found {
		return toTime(val)
	}

	return time.Time{}, ErrKeyNotFound
}

func (c *config) GetBool(key string) (bool, error) {
	val, found := c.handler.Get(key)
	if found {
		return toBool(val)
	}

	return false, ErrKeyNotFound
}

func (c *config) SetHandler(handler Handler) {
	c.handler = handler
}


func (c *config) Set(key string, value interface{}) error {
	return c.handler.Set(key, value)
}




















package config

import (
	"io"
	"time"
)

type Configer interface {
	Getter
	TypeGetter
	Setter
}

type config struct {
	handler Handler
}

func New() Configer {
	return nil
}

func (c *config) Load(reader io.Reader) error {
	return nil
}

func (c *config) Get(key string) (interface{}, bool) {
	return c.handler.Get(key)
}

func (c *config) GetInt(key string) (int, error) {
	val, found := c.Get(key)
	if found {
		return toInt(val)
	}

	return 0, ErrKeyNotFound
}

func (c *config) GetString(key string) (string, error) {
	val, found := c.Get(key)
	if found {
		return toString(val)
	}

	return "", ErrKeyNotFound
}

func (c *config) GetTime(key string) (time.Time, error) {
	val, found := c.Get(key)
	if found {
		return toTime(val)
	}

	return time.Time{}, ErrKeyNotFound
}

func (c *config) GetBool(key string) (bool, error) {
	val, found := c.Get(key)
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

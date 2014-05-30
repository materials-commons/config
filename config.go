package config

import (
	"time"
)

// A Configer is a configuration object that can store and retrieve key/value pairs.
type Configer interface {
	Initer
	Getter
	TypeGetter
	Setter
	SetHandler(handler Handler)
	SetHandlerInit(handler Handler) error
}

// config is a private type for storing configuration information.
type config struct {
	handler Handler
}

// New creates a new Configer instance that uses the specified Handler for
// key/value retrieval and storage.
func New(handler Handler) Configer {
	return &config{handler: handler}
}

// Init initializes the Configer. It should be called before retrieving
// or setting keys.
func (c *config) Init() error {
	return c.handler.Init()
}

// Get returns the value for a key. It can return any value type.
func (c *config) Get(key string) (interface{}, bool) {
	return c.handler.Get(key)
}

// GetInt returns an integer value for a key. See TypeGetter interface for
// error codes.
func (c *config) GetInt(key string) (int, error) {
	val, found := c.Get(key)
	if found {
		return ToInt(val)
	}

	return 0, ErrKeyNotFound
}

// GetString returns an string value for a key. See TypeGetter interface for
// error codes.
func (c *config) GetString(key string) (string, error) {
	val, found := c.Get(key)
	if found {
		return ToString(val)
	}

	return "", ErrKeyNotFound
}

// GetTime returns an time.Time value for a key. See TypeGetter interface for
// error codes.
func (c *config) GetTime(key string) (time.Time, error) {
	val, found := c.Get(key)
	if found {
		return ToTime(val)
	}

	return time.Time{}, ErrKeyNotFound
}

// GetBool returns an bool value for a key. See TypeGetter interface for
// error codes.
func (c *config) GetBool(key string) (bool, error) {
	val, found := c.Get(key)
	if found {
		return ToBool(val)
	}

	return false, ErrKeyNotFound
}

// SetHandler changes the handler for a Configer. If this method is called
// then you must call Init before accessing any of the keys.
func (c *config) SetHandler(handler Handler) {
	c.handler = handler
}

// SetHandlerInit changes the handler for a Configer. It also immediately calls
// Init and returns the error from this call.
func (c *config) SetHandlerInit(handler Handler) error {
	c.handler = handler
	return c.Init()
}

// Set sets key to value. See Setter interface for error codes.
func (c *config) Set(key string, value interface{}) error {
	return c.handler.Set(key, value)
}











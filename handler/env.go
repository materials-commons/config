package handler

import (
	"github.com/materials-commons/config"
	"os"
	"strings"
)

type envHandler struct{}

// EnvHandler returns a Handler that access keys that are environment variables.
func EnvHandler() config.Handler {
	return &envHandler{}
}

// Init initializes access to the environment.
func (h *envHandler) Init() error {
	return nil
}

// Get retrieves a environment variable. It assumes all keys are upper case.
// It will uppercase the key before attempting to retrieve its value.
func (h *envHandler) Get(key string) (interface{}, bool) {
	ukey := strings.ToUpper(key)
	val := os.Getenv(ukey)
	if val == "" {
		return val, false
	}
	return val, true
}

// Set sets an environment variable. It assumes all keys are upper case, and that
// values must be stored as strings. It will uppercase the key, and convert the
// value to a string before it attempts to store it. If the value cannot be
// converted to a string it returns ErrBadType.
func (h *envHandler) Set(key string, value interface{}) error {
	ukey := strings.ToUpper(key)
	sval, err := config.ToString(value)
	if err != nil {
		return config.ErrBadType
	}

	err = os.Setenv(ukey, sval)
	if err != nil {
		return config.ErrKeyNotSet
	}

	return nil
}
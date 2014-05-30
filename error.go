package config

import (
	"errors"
)

var (
	// ErrBadType conversion of value to specified type failed.
	ErrBadType = errors.New("bad type conversion")

	// ErrKeyNotFound specified key not found.
	ErrKeyNotFound = errors.New("key not found")
)

package config

import (
	"time"
)

// A Getter retrieves values for keys. It returns true if
// a value was found, and false otherwise.
type Getter interface {
	Get(key string) (interface{}, bool)
}

// A TypeGetter performs type safe conversion and retrieval of key values. The routines
// have return ErrBadType if the value doesn't match or can't be converted to the
// getter method called. ErrKeyNotFound is returned if the specified key is not found.
type TypeGetter interface {
	GetInt(key string) (int, error)
	GetString(key string) (string, error)
	GetTime(key string) (time.Time, error)
	GetBool(key string) (bool, error)
}

// A Setter stores a value for a key. It returns nil on success and ErrKeyNotSet on failure.
type Setter interface {
	Set(key string, value interface{}) error
}

// A Loader loads values into out.
type Loader interface {
	Load(out interface{}) error
}

// A Initer initializes a data structure.
type Initer interface {
	Init() error
}

// A Handler defines the interface to different types of stores. The Init method should
// always be called before attempting to get or set values for a handler.
type Handler interface {
	Initer
	Getter
	Setter
}

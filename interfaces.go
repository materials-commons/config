package config

import (
	"io"
	"time"
)

type Getter interface {
	Get(key string) (interface{}, bool)
}

type TypeGetter interface {
	GetInt(key string) (int, error)
	GetString(key string) (string, error)
	GetTime(key string) (time.Time, error)
	GetBool(key string) (bool, error)
}

type Setter interface {
	Set(key string, value interface{}) error
}

type Loader interface {
	Load(out interface{}) error
}

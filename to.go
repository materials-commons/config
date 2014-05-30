package config

import (
	"github.com/spf13/cast"
	"strconv"
	"time"
)

// toTime converts in to a time.Time value.
func toTime(in interface{}) (time.Time, error) {
	switch val := in.(type) {
	case int64:
		return time.Unix(val, 0), nil
	default:
		t, err := cast.ToTimeE(val)
		if err != nil {
			return t, ErrBadType
		}
		return t, nil
	}
}

// toBool converts in to a bool value.
func toBool(in interface{}) (bool, error) {
	switch val := in.(type) {
	case string:
		sval, err := strconv.ParseBool(val)
		if err != nil {
			return false, ErrBadType
		}
		return sval, nil
	default:
		value, err := cast.ToBoolE(val)
		if err != nil {
			return false, ErrBadType
		}
		return value, nil
	}
}

// toInt64 converts in to a int64 value.
func toInt(in interface{}) (int, error) {
	if val, err := cast.ToIntE(in); err == nil {
		return val, nil
	}
	return 0, ErrBadType
}

// toString converts in to a string value.
func toString(in interface{}) (string, error) {
	if val, err := cast.ToStringE(in); err == nil {
		return val, nil
	}
	return "", ErrBadType
}

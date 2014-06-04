package config

import (
	"strings"
)

// A configuration event
type Event int

const (
	// Get key event
	GET Event = iota

	// Set key event
	SET

	// Convert to int event
	TOINT

	// Convert to string event
	TOSTRING

	// Convert to bool event
	TOBOOL

	// Convert to time event
	TOTIME

	// Swap handler event
	SWAP

	// Initialize handler event
	INIT

	// Unknown event
	UNKNOWN
)

// tostr maps events to their string representation.
var tostr = map[Event]string{
	GET:      "GET",
	SET:      "SET",
	TOINT:    "TOINT",
	TOSTRING: "TOSTRING",
	TOBOOL:   "TOBOOL",
	TOTIME:   "TOTIME",
	SWAP:     "SWAP",
	INIT:     "INIT",
}

// String turns an event into a string.
func (e Event) String() string {
	if val, found := tostr[e]; found {
		return val
	}

	return "UNKNOWN"
}

// toevent maps strings to events
var toevent = map[string]Event{
	"GET":      GET,
	"SET":      SET,
	"TOINT":    TOINT,
	"TOSTRING": TOSTRING,
	"TOBOOL":   TOBOOL,
	"TOTIME":   TOTIME,
	"SWAP":     SWAP,
	"INIT":     INIT,
	"UNKNOWN":  UNKNOWN,
}

// ToEvent maps a string to an event. The match on string is not
// case sensitive.
func ToEvent(str string) Event {
	if val, found := toevent[strings.ToUpper(str)]; found {
		return val
	}

	return UNKNOWN
}

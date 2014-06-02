package handler

import (
	"github.com/materials-commons/config"
	"os"
	"testing"
)

func TestEnvHandler(t *testing.T) {
	var val interface{}
	h := Env()
	if err := h.Init(); err != nil {
		t.Fatalf("Init failed: %s", err)
	}

	// Get Non existent key
	os.Setenv("ENV_TEST", "") // make sure key is not set
	_, err := h.Get("ENV_TEST")
	if err != config.ErrKeyNotFound {
		t.Fatalf("Looked up of nonexistent key should have returned ErrKeyNotFound, instead: %s", err)
	}

	// Try non-existent lower case.
	_, err = h.Get("env_test")
	if err != config.ErrKeyNotFound {
		t.Fatalf("Looked up of nonexistent key should have returned ErrKeyNotFound, instead: %s", err)
	}

	// Set Key
	if err = h.Set("ENV_TEST", "TEST1"); err != nil {
		t.Fatalf("Setting key failed: %s", err)
	}

	if val, err = h.Get("ENV_TEST"); err != nil {
		t.Fatalf("Failed looking up key ENV_TEST: %s", err)
	}

	sval := val.(string)
	if sval != "TEST1" {
		t.Fatalf("ENV_TEST should be equal to 'TEST1', got %s", sval)
	}

	// Set key lower key, look up both ways
	if err = h.Set("env_test2", "TEST2"); err != nil {
		t.Fatalf("Setting key failed: %s", err)
	}

	if val, err = h.Get("ENV_TEST2"); err != nil {
		t.Fatalf("Failed looking up key ENV_TEST2: %s", err)
	}

	sval = val.(string)
	if sval != "TEST2" {
		t.Fatalf("ENV_TEST2 should be equal to 'TEST2', got %s", sval)
	}

	if val, err = h.Get("env_test2"); err != nil {
		t.Fatalf("Failed looking up key env_test2: %s", err)
	}

	sval = val.(string)
	if sval != "TEST2" {
		t.Fatalf("env_test2 should be equal to 'TEST2', got %s", sval)
	}

	// Make sure that Args is false
	if h.Args() {
		t.Fatalf("EnvHandler Args should have returned false.")
	}

	// Make sure that calls with extra args fail.
	if _, err := h.Get("ENV_TEST2", "EXTRA_ARG"); err != config.ErrArgsNotSupported {
		t.Fatalf("Get with extra args returned wrong err: %s", err)
	}

	if err := h.Set("ENV_TEST2", "BLAH", "EXTRA_ARG"); err != config.ErrArgsNotSupported {
		t.Fatalf("Set with extra args returned wrong err: %s", err)
	}

	// Make sure that the value wasn't set
	if val, err = h.Get("ENV_TEST2"); err != nil {
		t.Fatalf("Failed looking up key ENV_TEST2: %s", err)
	}

	if val == "BLAH" {
		t.Fatalf("Set with extra args set the value when it shouldn't have")
	}
}

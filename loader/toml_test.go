package loader

import (
	"bytes"
	"fmt"
	"testing"
)

var tomlTestData = []byte(`
[test]
stringkey = "stringvalue"
boolkey = true
intkey = 123
`)

func TestTOMLLoader(t *testing.T) {
	var (
		vals  map[string]interface{}
		tvals map[string]interface{}
		val   interface{}
		found bool
	)

	l := TOML(bytes.NewReader(tomlTestData))
	err := l.Load(&vals)
	if err != nil {
		t.Fatalf("failed loading TOML: %s", err)
	}

	fmt.Printf("vals = %#v\n", vals)

	if tvals, found = vals["test"]; !found {
		t.Fatalf("Failed to lookup key 'test'")
	}

	if val, found = tvals["stringkey"]; !found {
		t.Fatalf("Failed to lookup key 'stringkey'")
	}

	sval := val.(string)
	if sval != "stringvalue" {
		t.Fatalf("Unexpected value for sval '%s', expected 'stringvalue'", sval)
	}

	var _ = val
	var _ = found
}

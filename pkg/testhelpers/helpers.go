package testhelpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// ReadJSON read json from file and set object
func ReadJSON[T any](t *testing.T, file string) T {
	var obj T
	b, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		t.Fatal("failed to read file")
	}

	err = json.Unmarshal(b, &obj)
	if err != nil {
		t.Fatal("failed to unmarshal object")
	}
	return obj
}

// Read read json from file and return bytes
func Read(t *testing.T, file string) []byte {
	b, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		t.Fatalf("failed to read file %v\n", err)
	}
	return b
}

// CreateJSON creates a json file using the given object
func CreateJSON(t *testing.T, file string, object any) {
	b, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		t.Fatal("failed to create file")
	}

	/* #nosec */
	err = os.WriteFile(file, b, 0644)
	if err != nil {
		t.Fatal("failed to write json file")
	}
}

// RequestPost is a helper to create a http request post
func RequestPost(t *testing.T, httpMethod, file string) *http.Request {
	bReq := Read(t, file)
	req, err := http.NewRequest("POST", "http://localhost", bytes.NewBuffer(bReq))
	if err != nil {
		t.Fatal(err)
	}
	return req
}

// Compare it compares two objects and check if they are equals
func Compare(t *testing.T, description string, exp, got any, opts ...cmp.Option) {
	d := cmp.Diff(exp, got, opts...)
	if len(d) > 0 {
		t.Fatalf("test [%s] compare description [%s] mismatch (-want +got):\n%s", t.Name(), description, d)
	}
}

// CompareWithFile it compares two objects using expected file to check if they are equals
func CompareWithFile[T any](t *testing.T, description, expFile string, got T, opts ...cmp.Option) {
	exp := ReadJSON[T](t, expFile)
	Compare(t, description, exp, got)
}

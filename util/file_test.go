package util

import (
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	f, err := os.CreateTemp("", "exists_test")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	name := f.Name()
	f.Close()
	if !Exists(name) {
		t.Errorf("expected Exists(%s) to return true", name)
	}
	os.Remove(name)
	if Exists(name) {
		t.Errorf("expected Exists(%s) to return false after removal", name)
	}
}

package gifdecode

import (
	"os"
	"path/filepath"
	"testing"
)

func readFixture(t testing.TB, name string) []byte {
	t.Helper()
	path := filepath.Join("testdata", name)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	return data
}

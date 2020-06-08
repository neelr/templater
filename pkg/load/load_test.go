package load

import (
	"os"
	"testing"
)

func TestCommand(t *testing.T) {
	err := Command("test")
	if err != nil {
		t.Error(err)
	}
	os.Remove("create.go")
	os.Remove("create_test.go")
	os.Remove("debug.test")
}

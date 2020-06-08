package delete

import "testing"

func TestCommand(t *testing.T) {
	err := Command("test")
	if err != nil {
		t.Error(err)
	}
}

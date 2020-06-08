package list

import "testing"

func TestCommand(t *testing.T) {
	err := Command()
	if err != nil {
		t.Error(err)
	}
}

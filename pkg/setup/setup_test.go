package setup

import "testing"

func TestCommand(t *testing.T) {
	err := Configs()
	if err != nil {
		t.Error(err)
	}
}

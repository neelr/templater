package upload

import (
	"testing"

	"github.com/neelr/templater/config/settings"
)

func TestCommand(t *testing.T) {
	settings.InitSettings()

	err := Command("test")
	if err != nil {
		t.Error(err)
	}
}

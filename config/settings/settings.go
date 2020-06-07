package settings

import (
	"os"
)

// InitSettings initialize the settings for the app
func InitSettings() {

	// Which url to make server requests to
	os.Setenv("URL", "https://templater-api--hacker22.repl.co")

	// Firebase URL
	os.Setenv("FIREBASE", "templater-9289d.appspot.com")
}

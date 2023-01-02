package settings

import (
	"os"
)

// InitSettings initialize the settings for the app
func InitSettings() {

	// Which url to make server requests to
	os.Setenv("URL", "https://plate.neelr.dev")

	// Firebase URL
	os.Setenv("FIREBASE", "templater-9289d.appspot.com")

	// Github Client ID
	os.Setenv("GH_CLIENT_ID", "33449dcb2152190815b2")
}

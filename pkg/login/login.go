package login

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"

	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
)

func Command() error {
	setup.Configs()
	key, _ := randomHex(128)
	log.InformationPrint("Redirecting you to GitHub Oauth...")
	// Check if theres a key already in .config
	startKey, err := ioutil.ReadFile(path.Join(os.Getenv("PLATE_DIR"), ".config"))
	if err != nil {
		openbrowser("https://github.com/login/oauth/authorize?client_id=" + os.Getenv("GH_CLIENT_ID") + "&state=" + key)
	} else {
		// If a key, send it over to delete
		openbrowser("https://github.com/login/oauth/authorize?client_id=" + os.Getenv("GH_CLIENT_ID") + "&state=" + key + "|" + string(startKey))
	}

	return ioutil.WriteFile(path.Join(os.Getenv("PLATE_DIR"), ".config"), []byte(key), 0644)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.ErrorPrint(err.Error())
	}

}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

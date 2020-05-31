package setup

import (
	"fmt"
	"os"
	"os/user"

<<<<<<< HEAD
	log "github.com/neelr/templater/pkg/logs"
=======
	log "github.com/neelr/plate/pkg/logs"
>>>>>>> a850638d1d5a25cf0f5818826143b206e99a6ea0
)

// Configs sets up the config folder
func Configs() {
	dir := os.Getenv("PLATE_DIR")
	if os.Getenv("PLATE_DIR") == "" {
		usr, err := user.Current()

		if err != nil {
			fmt.Println("Home Directory Not Found")
			return
		}

		dir = usr.HomeDir + "/.plate"
		os.Setenv("PLATE_DIR", dir)
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.ErrorPrint("Config folder not found... Creating at " + dir)
		os.Mkdir(dir, os.ModePerm)
	}

}

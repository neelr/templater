package create

import (
	"os"
	"path"

<<<<<<< HEAD
	log "github.com/neelr/templater/pkg/logs"
	setup "github.com/neelr/templater/pkg/setup"
=======
	log "github.com/neelr/plate/pkg/logs"
	setup "github.com/neelr/plate/pkg/setup"
>>>>>>> a850638d1d5a25cf0f5818826143b206e99a6ea0
	"github.com/otiai10/copy"
)

// RunCommand runs a command
func Command(name string) {
	setup.Configs()

	configFolder := path.Join(os.Getenv("PLATE_DIR"), name)
	currentFolder, _ := os.Getwd()

	log.Loading.Suffix = log.Information("Transferring Files....")
	log.Loading.Start()
	copy.Copy(currentFolder, configFolder)
	log.Loading.Stop()
	log.InformationPrint("Saved current directory as \"" + name + "\"!")
}

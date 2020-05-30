package create

import (
	"os"
	"path"

	log "github.com/neelr/plate/pkg/logs"
	setup "github.com/neelr/plate/pkg/setup"
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

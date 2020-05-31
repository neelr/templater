package load

import (
	"os"
	"path"

	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
	"github.com/otiai10/copy"
)

func Command(name string) {
	setup.Configs()

	configFolder := path.Join(os.Getenv("PLATE_DIR"), name)
	currentFolder, _ := os.Getwd()

	log.Loading.Suffix = log.Information("Transferring Files....")
	log.Loading.Start()
	copy.Copy(configFolder, currentFolder)
	log.Loading.Stop()

	log.InformationPrint("Transferred template to current directory!")
}

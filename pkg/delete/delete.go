package delete

import (
	"os"
	"path"

	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
)

func Command(name string) {
	setup.Configs()

	configFolder := path.Join(os.Getenv("PLATE_DIR"), name)
	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		log.ErrorPrint("Template not found!")
	} else {
		log.Loading.Suffix = log.Information("Deleting Files....")
		log.Loading.Start()
		os.RemoveAll(configFolder)
		log.Loading.Stop()

		log.InformationPrint("Deleted \"" + name + "\"!")
	}
}

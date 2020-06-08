package delete

import (
	"errors"
	"os"
	"path"

	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
)

func Command(name string) error {
	setup.Configs()

	configFolder := path.Join(os.Getenv("PLATE_DIR"), name)
	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		log.ErrorPrint("Template not found!")
		return errors.New("template not found")
	}
	log.Loading.Suffix = log.Information("Deleting Files....")
	log.Loading.Start()
	os.RemoveAll(configFolder)
	log.Loading.Stop()

	log.InformationPrint("Deleted \"" + name + "\"!")
	return nil
}

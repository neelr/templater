package list

import (
	"io/ioutil"
	"os"
	"strconv"

	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
)

func Command() {
	setup.Configs()

	templates, _ := ioutil.ReadDir(os.Getenv("PLATE_DIR"))

	log.InformationPrint("All Templates")
	log.InformationPrint("------------------")
	for _, f := range templates {
		if f.IsDir() {
			log.NormalPrint(f.Name() + " - " + strconv.FormatInt(f.Size(), 10) + " bytes")
		}
	}
}

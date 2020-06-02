package removeserver

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/neelr/templater/pkg/login"
	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
)

func Command(name string) {
	setup.Configs()
	configFile := path.Join(os.Getenv("PLATE_DIR"), ".config")
	key, err := ioutil.ReadFile(configFile)
	if err != nil {
		login.Command()
		return
	}
	log.Loading.Suffix = log.Error(" Deleting the template from our servers...")
	log.Loading.Start()
	resp, err := http.Get("https://templater-api--hacker22.repl.co/api/templates/delete?key=" + string(key) + "&template=" + name)
	log.Loading.Stop()
	if err != nil {
		log.ErrorPrint(err.Error())
		return
	}
	if resp.StatusCode == 200 {
		log.InformationPrint("Successfully Deleted!")
		return
	}
	log.ErrorPrint("Template not found...")
}

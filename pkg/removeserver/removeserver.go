package removeserver

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/neelr/templater/pkg/login"
	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
)

func Command(name string) error {
	setup.Configs()
	configFile := path.Join(os.Getenv("PLATE_DIR"), ".config")
	key, err := ioutil.ReadFile(configFile)
	if err != nil {
		login.Command()
		return errors.New("no config file")
	}
	log.Loading.Suffix = log.Error(" Deleting the template from our servers...")
	log.Loading.Start()
	resp, err := http.Get(os.Getenv("URL") + "/api/template/delete?key=" + string(key) + "&template=" + name)
	log.Loading.Stop()
	if err != nil {
		log.ErrorPrint(err.Error())
		return err
	}
	if resp.StatusCode == 200 {
		log.InformationPrint("Successfully Deleted!")
		return nil
	}
	log.ErrorPrint("Template not found...")
	return errors.New("template not found")
}

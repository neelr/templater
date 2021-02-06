package list

import (
	"io/ioutil"
	"os"
	"fmt"
	"path"
	"path/filepath"

	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
)

func DirSize(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            size += info.Size()
        }
        return err
    })
    return size, err
}

func Command() error {
	setup.Configs()

	templates, err := ioutil.ReadDir(os.Getenv("PLATE_DIR"))
	if err != nil {
		return err
	}
	log.InformationPrint("All Templates")
	log.InformationPrint("------------------")
	for _, f := range templates {
		size, _ := DirSize(path.Join(os.Getenv("PLATE_DIR"),f.Name())); 
		if f.IsDir() {
			log.NormalPrint(f.Name() + " - " + fmt.Sprintf("%.2f",float32(size)/1000.0) + " KB")
		}
	}
	return nil
}

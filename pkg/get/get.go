package get

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
)

func Command(slug string) error {
	setup.Configs()

	tmpDir := path.Join(os.Getenv("PLATE_DIR"), "tmp.zip")

	log.Loading.Suffix = log.Information(" Getting template...")
	log.Loading.Start()
	status := downloadFile(tmpDir, os.Getenv("URL")+"/api/templates/"+slug+"/download")
	if status == 404 || status == 0 {
		log.Loading.Stop()
		log.ErrorPrint("Cannot find template!")
		return errors.New("no such template")
	}
	unzip(tmpDir, os.Getenv("PLATE_DIR"))
	os.Remove(tmpDir)
	log.InformationPrint("Loaded " + slug)
	return nil
}

func downloadFile(filepath string, url string) int {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return resp.StatusCode
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return resp.StatusCode
}
func unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

package upload

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/neelr/templater/pkg/login"
	log "github.com/neelr/templater/pkg/logs"
	"github.com/neelr/templater/pkg/setup"
)

func Command(name string) {
	client := &http.Client{}
	setup.Configs()
	configFile := path.Join(os.Getenv("PLATE_DIR"), ".config")
	if _, err := os.Stat(configFile); err != nil {
		login.Command()
	}
	if _, err := os.Stat(path.Join(os.Getenv("PLATE_DIR"), name)); err != nil {
		log.ErrorPrint("Template does not exist!")
		return
	}

	log.Loading.Suffix = log.Information(" Uploading the template....")
	log.Loading.Start()
	zipit(path.Join(os.Getenv("PLATE_DIR"), name), path.Join(os.Getenv("PLATE_DIR"), "tmp.zip"))

	key, _ := ioutil.ReadFile(configFile)
	readme, err := ioutil.ReadFile(path.Join(os.Getenv("PLATE_DIR"), name, "README.md"))
	if err != nil {
		readme = []byte("")
	}
	var allFiles []string
	err = filepath.Walk(path.Join(os.Getenv("PLATE_DIR"), name),
		func(file string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				allFiles = append(allFiles, strings.Replace(file, path.Join(os.Getenv("PLATE_DIR"), name), "", 1))
				return nil
			}
			return nil
		})
	marshalledFiles, _ := json.Marshal(allFiles)
	req, err := newfileUploadRequest("https://templater-api--hacker22.repl.co/api/upload", map[string]string{
		"name":   name,
		"README": string(readme),
		"files":  string(marshalledFiles),
		"state":  string(key),
	}, "zip", path.Join(os.Getenv("PLATE_DIR"), "tmp.zip"))
	r, err := client.Do(req)
	if err != nil {
		log.Loading.Stop()
		log.ErrorPrint("Couldn't connect to servers... Try Again Later!")
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	log.Loading.Stop()
	os.Remove(path.Join(os.Getenv("PLATE_DIR"), "tmp.zip"))
	if r.StatusCode == 200 {
		log.InformationPrint("Uploaded file to " + string(body))

	}
	log.ErrorPrint("Error " + r.Status + "! Make sure you are logged in! If this doesn't make sense, make an issue on https://github.com/neelr/templater!")
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	part, err := writer.CreateFormFile(paramName, fi.Name())
	io.Copy(part, file)
	file.Close()
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request, nil
}

func zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/neelr/templater/pkg/create"
	"github.com/neelr/templater/pkg/delete"
	"github.com/neelr/templater/pkg/get"
	"github.com/neelr/templater/pkg/list"
	"github.com/neelr/templater/pkg/load"
	"github.com/neelr/templater/pkg/login"
	log "github.com/neelr/templater/pkg/logs"
	removeserver "github.com/neelr/templater/pkg/removesever"
	"github.com/neelr/templater/pkg/upload"
)

func main() {
	helpText := `
		HELP:
		create {name} - Creates a template from the current directory and stores it in the name
		load {name} - Loads a template to the current directory
		delete {name} - Deletes a template with that name
		list - Lists all downloaded and created templates along with file size
		login - Logs you in to our server's for uploads
		upload {name} - Uploads the template to your cloud/account
		get {user}/{name} - Gets the template of the user and downloads to your templates
		deletefromserver {name} - Deletes the template that you uploaded before
	`
	if len(os.Args) < 2 {
		log.NormalPrint(helpText)
		return
	}

	if len(os.Args) >= 3 {
		switch os.Args[1] {
		case "create":
			create.Command(os.Args[2])
		case "load":
			load.Command(os.Args[2])
		case "delete":
			delete.Command(os.Args[2])
		case "upload":
			upload.Command(os.Args[2])
		case "deletefromserver":
			removeserver.Command(os.Args[2])
		case "get":
			get.Command(os.Args[2])
		case "version":
			log.NormalPrint("plate v1.0")
		default:
			log.NormalPrint(helpText)
		}
	} else if os.Args[1] == "list" {
		list.Command()
	} else if os.Args[1] == "login" {
		login.Command()
	} else {
		log.NormalPrint(helpText)
	}
}
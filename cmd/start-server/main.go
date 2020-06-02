package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/neelr/templater/pkg/logs"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	// Github OAUTH
	http.HandleFunc("/api/oauth", handleOauth)

	// Template uploads
	http.HandleFunc("/api/upload", upload)

	// Get template data
	http.HandleFunc("/api/templates/", getData)

	// Get user data
	http.HandleFunc("/api/user/", userData)

	// Get all users
	http.HandleFunc("/api/users", allUsers)

	log.InformationPrint("On port 3000!")
	http.ListenAndServe(":3001", nil)
}

func handleOauth(w http.ResponseWriter, r *http.Request) {
	// Setup Firebase creds
	ctx := context.Background()
	sa := option.WithCredentialsFile("./firebase_service.json")
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "templater-9289d",
	}, sa)

	// Create firestore instance
	client, err := app.Firestore(ctx)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer client.Close()

	// Take code and state params from github oauth
	code := r.URL.Query()["code"]
	state := r.URL.Query()["state"]

	if code == nil || state == nil {
		w.WriteHeader(401)
		return
	}

	// Get Oauth token
	resp, err := http.PostForm("https://github.com/login/oauth/access_token", url.Values{
		"code":          {code[0]},
		"client_id":     {os.Getenv("CLIENT_ID")},
		"client_secret": {os.Getenv("CLIENT_SECRET")},
	})
	if err != nil {
		log.ErrorPrint(err.Error())
	}
	defer resp.Body.Close()

	// read oauth token and save to firebase map
	body, err := ioutil.ReadAll(resp.Body)
	params, _ := url.ParseQuery(string(body))

	if params["access_token"] == nil {
		w.WriteHeader(401)
		return
	}
	// Check if the state includes a key to delete
	keys := strings.SplitN(state[0], "|", 2)
	key := keys[0]
	// If it has a key to delete, delete it
	if len(keys) == 2 {
		client.Collection("key2oauth").Doc("map").Update(ctx, []firestore.Update{
			{
				Path:  keys[1],
				Value: firestore.Delete,
			},
		})
	}

	client.Collection("key2oauth").Doc("map").Set(ctx, map[string]interface{}{
		key: params["access_token"][0],
	}, firestore.MergeAll)
	w.Write([]byte("Logged in! You can go back to terminal!"))
}

func upload(w http.ResponseWriter, r *http.Request) {
	// Setup Firestore and Storage
	httpClient := &http.Client{}
	ctx := context.Background()
	sa := option.WithCredentialsFile("./firebase_service.json")
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "templater-9289d",
	}, sa)
	client, err := app.Firestore(ctx)
	storageClient, err := app.Storage(ctx)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer client.Close()

	// read multipart
	readForm, err := r.MultipartReader()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	// Set up variables to be filled
	name := ""
	zipUUID := uuid.New().String()
	readme := ""

	authed := false
	for {
		part, err := readForm.NextPart()
		if err == io.EOF {
			break
		}
		if part.FormName() == "zip" {

			// upload zipped template if the authentication state came before
			if !authed {
				w.WriteHeader(401)
				return
			}

			// Setup Bucket
			bucket, err := storageClient.Bucket("templater-9289d.appspot.com")
			if err != nil {
				w.WriteHeader(400)
				return
			}
			wc := bucket.Object(zipUUID + ".zip").NewWriter(ctx)
			wc.ContentType = "application/zip"
			buf := new(bytes.Buffer)
			buf.ReadFrom(part)

			// Write zip to bucket
			if _, err := wc.Write(buf.Bytes()); err != nil {
				log.ErrorPrint(err.Error())
				return
			}
			if err := wc.Close(); err != nil {
				log.ErrorPrint(err.Error())
				return
			}
		} else if part.FormName() == "name" {
			// Save Name
			buf := new(bytes.Buffer)
			buf.ReadFrom(part)
			name = buf.String()
		} else if part.FormName() == "README" {
			// Save README string
			buf := new(bytes.Buffer)
			buf.ReadFrom(part)
			readme = buf.String()
		} else if part.FormName() == "state" {
			// AUTHORIZE STATE
			buf := new(bytes.Buffer)
			buf.ReadFrom(part)
			snap, err := client.Collection("key2oauth").Doc("map").Get(ctx)
			if err != nil || snap.Data()[buf.String()] == nil {
				w.WriteHeader(401)
				return
			}
			// Got oauth github, now request the login username
			req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
			req.Header.Add("Authorization", "token "+snap.Data()[buf.String()].(string))
			resp, err := httpClient.Do(req)
			if err != nil {
				w.WriteHeader(400)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.ErrorPrint(err.Error())
				w.WriteHeader(400)
				return
			}

			var githubResponse map[string]interface{}

			json.Unmarshal(body, &githubResponse)

			// Check if tenplate exists
			snap, err = client.Collection("Users").Doc(githubResponse["login"].(string)).Collection("templates").Doc(name).Get(ctx)
			if err == nil {
				// If exists, delete the zip file in storage so we can replace later on

				// Setup Bucket
				bucket, err := storageClient.Bucket("templater-9289d.appspot.com")
				if err != nil {
					w.WriteHeader(400)
					return
				}
				err = bucket.Object(snap.Data()["id"].(string) + ".zip").Delete(ctx)
			}

			// Got and parsed login username, now store new template document in firebase as well as reference to storage file
			_, err = client.Collection("Users").Doc(githubResponse["login"].(string)).Collection("templates").Doc(name).Set(ctx, map[string]string{
				"id":     zipUUID,
				"README": readme,
			})

			// May as well update the bio while your at it for a speedy website
			_, err = client.Collection("Users").Doc(githubResponse["login"].(string)).Set(ctx, map[string]string{
				"avatar": githubResponse["avatar_url"].(string),
				"bio":    githubResponse["bio"].(string),
				"url":    githubResponse["html_url"].(string),
			})

			if err != nil {
				log.ErrorPrint(err.Error())
				w.WriteHeader(400)
				return
			}
			w.Write([]byte(githubResponse["login"].(string) + "/" + name))
			authed = true
		}
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	// Setup Firestore and Storage
	ctx := context.Background()
	sa := option.WithCredentialsFile("./firebase_service.json")
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "templater-9289d",
	}, sa)
	client, err := app.Firestore(ctx)
	storageClient, err := app.Storage(ctx)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer client.Close()

	// Split Path
	pathArray := strings.SplitN(r.URL.Path, "/", 6)
	// If it doesn't have at least 5 elements (no name or template)
	if !(len(pathArray) >= 5) {
		w.WriteHeader(400)
		return
	}
	name := pathArray[3]
	templateName := pathArray[4]

	// retrieve template with name and template giving
	templateDesc, err := client.Collection("Users").Doc(name).Collection("templates").Doc(templateName).Get(ctx)

	// If err, that means not there so 404
	if err != nil {
		w.WriteHeader(404)
		return
	}

	// if it wants to download
	if len(pathArray) == 6 && pathArray[5] == "download" {

		// Get storage bucket and copy from there to response
		bucket, _ := storageClient.Bucket("templater-9289d.appspot.com")
		rc, _ := bucket.Object(templateDesc.Data()["id"].(string) + ".zip").NewReader(ctx)
		w.Header().Set("Content-Disposition", "attachment; filename="+templateName+".zip")
		w.Header().Add("Content-Type", "application/zip")
		io.Copy(w, rc)
	} else {
		// If want metadata, then give the data of description in JSON
		w.Header().Add("Content-Type", "application/json")
		packetJSON, _ := json.Marshal(templateDesc.Data())
		w.Write(packetJSON)
	}
}

func userData(w http.ResponseWriter, r *http.Request) {
	// Setup Firestore and Storage
	ctx := context.Background()
	sa := option.WithCredentialsFile("./firebase_service.json")
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "templater-9289d",
	}, sa)
	client, err := app.Firestore(ctx)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer client.Close()

	// Parse path
	pathArray := strings.SplitN(r.URL.Path, "/", 4)

	// If it doesn't contain a user
	if !(len(pathArray) >= 4) {
		w.WriteHeader(400)
		return
	}
	name := pathArray[3]

	// Get user data
	user, err := client.Collection("Users").Doc(name).Get(ctx)

	// If err, then no user
	if err != nil {
		w.WriteHeader(404)
		return
	}

	// Get all templates from user
	templatesDoc := client.Collection("Users").Doc(name).Collection("templates").Documents(ctx)

	// Iterate over collection of templates to get each template name
	var templates []string
	for {
		doc, err := templatesDoc.Next()
		if err == iterator.Done {
			break
		}
		paths := strings.SplitN(doc.Ref.Path, "/", 9)
		templates = append(templates, paths[8])
	}

	// Send over user data + templates we just collected
	w.Header().Add("Content-Type", "application/json")
	packet := user.Data()
	packet["templates"] = templates
	packetJSON, _ := json.Marshal(packet)
	w.Write(packetJSON)
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	// Setup Firestore and Storage
	ctx := context.Background()
	sa := option.WithCredentialsFile("./firebase_service.json")
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "templater-9289d",
	}, sa)
	client, err := app.Firestore(ctx)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer client.Close()

	// Get all users
	usersDocs := client.Collection("Users").Documents(ctx)

	// Iterate over collection of users to get each user name
	var users []string
	for {
		doc, err := usersDocs.Next()
		if err == iterator.Done {
			break
		}
		paths := strings.SplitN(doc.Ref.Path, "/", 7)
		users = append(users, paths[6])
	}

	// Send array of users
	w.Header().Add("Content-Type", "application/json")
	packet, _ := json.Marshal(users)
	w.Write(packet)
}

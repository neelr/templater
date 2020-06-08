# templater

<p align="center">
    <a href="https://templaterx.now.sh" alt="Templates">
        <img src="https://img.shields.io/endpoint?url=https://templater-api.hacker22.repl.co/api/badges/templates&label=Templates" /></a>
  <a href="https://templaterx.now.sh" alt="Server">
        <img src="https://img.shields.io/endpoint?url=https://templater-api.hacker22.repl.co/api/badges&label=Server%20Status" /></a>
  <a href="https://opensource.org/licenses/MIT" alt="LICENSE">
        <img src="https://badgen.net/github/license/neelr/templater" /></a>
    <a href="https://github.com/neelr/templater/commits/master" alt="Commits">
        <img src="https://badgen.net/github/last-commit/neelr/templater" /></a>
    <a href="https://github.com/neelr/templater/issues" alt="Closed Issues">
        <img src="https://badgen.net/github/closed-issues/neelr/templater" /></a>
     <a href="https://github.com/neelr/templater/issues" alt="Open Issues">
        <img src="https://badgen.net/github/open-issues/neelr/templater" /></a>
    <a href="https://github.com/neelr/templater/actions" alt="Actions">
        <img src="https://badgen.net/github/checks/neelr/templater" /></a>
      <a href="https://github.com/neelr/templater/actions" alt="Actions">
        <img src="https://badgen.net/github/checks/neelr/templater" /></a>
  <a href="https://github.com/neelr/templater/releases" alt="Release">
        <img src="https://badgen.net/github/release/neelr/templater" /></a>
  
</p>

A cli to create and share code templates and structure! This makes life super easy, so you don't have to create the React folder structure, and can create production ready structure in seconds! Also the best part about this, is since it has integration with GitHub, you can upload and download other people's structures, so the possibilities are endless!

### Installation

1. `go get -u github.com/neelr/templater/cmd/plate` Get the package from github

2. `go install github.com/neelr/templater/cmd/plate` Install the package!

### Docs

##### `plate create name`

Creates a template using the files in the current directory, and calls it "name"!

##### `plate load name`

Loads the specified temnplate into the current directory

##### `plate delete name`

Deletes the template "name"

##### `plate list`

Lists all installed templates

##### `plate login`

Logs you in so you can upload your templates so others can download them!

##### `plate upload name`

After you've logged in, you can upload a plate so people can see it! It'll give you a slug in the format `githubUsername/templateName` which others can use to download your template!

##### `plate deletefromserver name`

Delete the template you uploaded in the past!

##### `plate get githubUsername/templateName`

Install the template someone uploaded! You can load it anytime you want!

## TODO:

-   [x] Create a Nextjs UI so people can view people's templates
-   [x] Create a landing page ~along with a logo~
-   [x] Add Badges

Open to contributers! Just open up an issue/PR!

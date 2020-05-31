# templater
A cli to create and share code templates and structure! This makes life super easy, so you don't have to create the React folder structure, and can create production ready structure in seconds! Also the best part about this, is since it has integration with GitHub, you can upload and download other people's structures, so the possibilities are endless!

### Installation


1. `go get -u github.com/neelr/templater/cmd/plate` Get the package from github

2. `go install ithub.com/neelr/templater/cmd/plate` Install the package!

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

##### `plate get githubUsername/templateName`

Install the template someone uploaded! You can load it anytime you want!


## TODO:
1. Create a Nextjs UI so people can view people's templates
2. Create a landing page along with a logo

Open to contributers! Just open up an issue/PR!

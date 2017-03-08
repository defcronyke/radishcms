RadishCMS

Windows Installation:

Install the Go version of the Google App Engine SDK:
	https://cloud.google.com/appengine/docs/standard/go/download
	
Install the Google Cloud SDK:
	https://cloud.google.com/sdk/

Install Golang from:
	https://golang.org/dl/
	
Set a $GOPATH:
	GOPATH=$HOME/go
	
Install npm and nodejs:
	https://nodejs.org/en/download/
	
Install Yeoman and FountainJS framework:
	npm install -g yo generator-fountain-webapp

Initialize a new FountainJS project in the static folder:	
	yo fountain-webapp --framework=angular2 --modules=webpack --js=babel --css=scss --ci=travis
	At the prompts, choose:
		A working landing page
		@angular/router
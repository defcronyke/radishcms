RadishCMS

Install on Windows:  

Install the Go version of the Google App Engine SDK:  
	https://cloud.google.com/appengine/docs/standard/go/download  

Install the Google Cloud SDK:  
	https://cloud.google.com/sdk/  

Install Python 2.7:  
	https://www.python.org/downloads/  

Install Golang from:  
	https://golang.org/dl/  

Set a $GOPATH env var to some new folder in your home dir.  

Install npm and nodejs:  
	https://nodejs.org/en/download/  

Install Yeoman, FountainJS framework, and Gulp:  
	npm install -g yo generator-fountain-webapp gulp-cli  

Initialize a new FountainJS project in the static folder:  
	yo fountain-webapp --framework=angular2 --modules=webpack --js=babel --css=scss --ci=travis  
	At the prompts, choose:  
		A working landing page  
		@angular/router  

Make a new Google Cloud project for your website, and initialize  
App Engine for the project.  

Install some backend dependencies:  
	go get -u github.com/gorilla/mux  
	go get -u github.com/gorilla/context  
	go get -u google.golang.org/appengine/datastore
	go get -u github.com/rs/cors

Run the dev server for the backend. In the default folder:  
	goapp serve app.yaml  

Run the dev server for the frontend. In the default/static folder:  
	gulp serve  

*** Develop your site ***  

When a version of the site is ready to be published, in default/static:  
	gulp build  

Then in the default folder:  
	goapp deploy app.yaml  

# Go Project Template
Parses all files residing in [files](./files/) and puts them where told.

## Usage
To put the files into the current work dir:
```bash
go run main.go
```

To put the files to a custom path:
```bash
go run main.go -targetDir=/home/user/go/src/github.com/seibert-media/new-project
```

After running the code you might want to call `make deps`, `dep init` and `git init`.

## Current Setup
The current setup generates our own project template using the following variables:
```.env
export TEMP_ORG=//S/M
export TEMP_NAME=new-project
export TEMP_KEY=new-project
export TEMP_GIT_ORG=seibert-media
export TEMP_GIT_HOST=github.com
export TEMP_DOCKER_HOST=quay.io
export TEMP_DOCKER_ORG=seibertmedia
export TEMP_DOCKER_IMAGE=${TEMP_NAME}
export TEMP_DOCKER_MAINTAINER="//SEIBERT/MEDIA GmbH <docker@seibert-media.net>"
```

## License
Do what you want with it :-)
Issues, PR's and Feedback are highly welcome.
Some of the stuff we use in our project templates is based on the great work at github.com/kolide/kit
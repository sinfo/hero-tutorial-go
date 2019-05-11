# Hero REST API in Golang

## Setup

```
# check if you have the env var GOPATH set
$ echo $GOPATH

# if this outputs nothing, then set it to wherever you want
$ export GOPATH=/some/path

# create the directory for this project
$ mkdir -p $GOPATH/src/github.com/sinfo

# go inside that dir
$ cd $GOPATH/src/github/sinfo

# clone this repository
$ git clone git@github.com:sinfo/hero-tutorial-go.git

# install dependencies
$ make deps
```

## Build and run

This required an instance of MongoDB running.

Access this on `http://localhost:8080`

```
$ make run
```

## Docker

This requires docker to be installed and running.

### Test

This is what will be run on the dockerhub.

```
$ make docker-test
```

### Run

This required an instance of MongoDB running.

Access this on `http://localhost` (`http://localhost/documentation/` for the swagger documentation)

```
$ make docker
```

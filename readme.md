## go

### setup

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

### build and run

**Note:** don't forget to spin up a mongodb instance

```
# run
$ make run
```

### docker

```
# test
$ make docker-test

# run
$ make docker-run
```

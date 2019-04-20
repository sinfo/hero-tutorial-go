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
```

### build and run

**Note:** don't forget to spin up a mongodb instance

```
# build
$ go build -o main *.go

# test
$ go -c test ./routes
$ ./routes.test

# run
$ ./main
```

### development

#### lint ([revive](https://github.com/mgechev/revive))
```
$ revive -exclude vendor/... -formatter friendly ./...
```

### docker

```
# build and test
$ docker-compose -f docker-compose-test.yml -p ci build
$ docker-compose -f docker-compose-test.yml -p ci up --abort-on-container-exit
```

### dependencies

- [gorilla/mux](https://github.com/gorilla/mux)
- [dep](https://github.com/golang/dep)
- [mgo](https://godoc.org/github.com/globalsign/mgo)

## go

### setup

```
# check if you have the env var GOPATH set
$ echo $GOPATH

# if this outputs nothing, then set it to wherever you want
$ export $GOPATH=/some/path

# create the directory for this project
$ mkdir -p $GOPATH/src/github.com/sinfo

# go inside that dir
$ cd $GOPATH/src/github/sinfo

# clone this repository
$ git clone git@github.com:sinfo/hero-tutorial-go.git
```

### build and run
```
# build
$ go build -o main *.go

# test
$ go test ./...

# run linter
$ golint ./...

#run
$ ./main
```

## docker

```
# build and run
$ docker-compose up --build
```

## dependencies

- [gorilla/mux](https://github.com/gorilla/mux)
- [dep](https://github.com/golang/dep)
- [mgo](https://godoc.org/github.com/globalsign/mgo)

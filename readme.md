## build
```
docker build -t go-tutorial .
```

## run
```
docker run -it -p 3000:3000 --name some-mongo -d mongo:3.4.20-xenial go-tutorial
```

## development

See https://github.com/gravityblast/fresh

## dependencies

- [gorilla/mux](https://github.com/gorilla/mux)
- [dep](https://github.com/golang/dep)
- [mgo](https://godoc.org/github.com/globalsign/mgo)

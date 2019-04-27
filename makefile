BINDIR=./bin
SRCDIR=./src
BINARY_FILENAME=main

all: build test

build: swagger.yml src/*
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/$(BINARY_FILENAME) $(SRCDIR)/*.go
	go test -c $(SRCDIR)/routes -o $(BINDIR)/routes.test
	go test -c $(SRCDIR)/models -o $(BINDIR)/models.test
	go test -c $(SRCDIR)/server -o $(BINDIR)/server.test
	swagger generate spec -m -o swagger.yml -b ./src

test: build
	swagger validate ./swagger.yml
	chmod +x ./scripts/run_tests
	./scripts/run_tests

run: build
	$(BINDIR)/$(BINARY_FILENAME)

deps:
	go get -d -v ./...
	go get -u gotest.tools/assert
	go get -u golang.org/x/lint/golint
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

docker-build:
	docker-compose -f docker-compose.yml build
	docker-compose -f docker-compose.yml up --abort-on-container-exit

docker-test:
	docker-compose -f docker-compose-test.yml -p ci build
	docker-compose -f docker-compose-test.yml -p ci up --abort-on-container-exit

clean:
	go clean
	rm -rf data ./$(BINDIR) swagger.yml
BINDIR=./bin
SRCDIR=./src
BINARY_FILENAME=main

.PHONY: docker
all: build test

build: src/*
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/$(BINARY_FILENAME) $(SRCDIR)/*.go
	go test -c $(SRCDIR)/routes -o $(BINDIR)/routes.test
	go test -c $(SRCDIR)/models -o $(BINDIR)/models.test
	go test -c $(SRCDIR)/server -o $(BINDIR)/server.test
	swagger generate spec -m -o ./static/swagger.json -b ./src

test: build
	swagger validate ./static/swagger.json
	chmod +x ./scripts/run_tests
	./scripts/run_tests

run: build
	$(BINDIR)/$(BINARY_FILENAME)

deps: docker-deps
	go get -u gotest.tools/assert
	go get -u golang.org/x/lint/golint
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

# === Docker related ===
docker:
	docker-compose -f docker-compose.yml build
	docker-compose -f docker-compose.yml up --abort-on-container-exit

docker-run:
	$(BINDIR)/$(BINARY_FILENAME)

docker-deps:
	go get github.com/golang/dep/cmd/dep
	dep ensure -vendor-only

docker-build: docker-deps
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/$(BINARY_FILENAME) $(SRCDIR)/*.go

docker-test:
	docker-compose -f docker-compose.test.yml -p ci build
	docker-compose -f docker-compose.test.yml -p ci up --abort-on-container-exit

clean:
	go clean
	rm -rf data $(BINDIR) ./static/swagger.json
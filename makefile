BINDIR=./bin
SRCDIR=./src
BINARY_FILENAME=main

all: test build

build: *.go
	mkdir -p $(BINDIR)
	go build -o $(BINDIR)/$(BINARY) $(SRCDIR)/*.go
	go test -c ./routes -o $(BINDIR)/routes.test
	go test -c ./models -o $(BINDIR)/models.test
	go test -c ./server -o $(BINDIR)/server.test
	swagger generate spec -m -o swagger.yml

test: build
	swagger validate ./swagger.yml
	chmod +x ./scripts/run_tests
	./scripts/run_tests
run: build
	./$(BINDIR)/$(BINARY)

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
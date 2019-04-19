FROM golang:alpine

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

RUN apk add git
RUN apk add gcc
RUN apk add libc-dev

RUN mkdir -p "$GOPATH/src/github/sinfo/go-tutorial"
WORKDIR $GOPATH/src/github.com/sinfo/go-tutorial/
COPY . .

RUN go get -d -v ./...
RUN go get gotest.tools/assert
RUN go install -v ./...

RUN go build -o main *.go
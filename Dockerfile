FROM golang:1.12.1

RUN apt-get update; apt-get install curl; apt-get install git
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 \
      && chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/sinfo/go-tutorial
COPY . .

RUN dep ensure -vendor-only
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-tutorial"]

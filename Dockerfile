FROM golang:latest

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

RUN apt-get update
RUN apt-get install git gcc

RUN mkdir -p "$GOPATH/src/github/sinfo/go-tutorial"
WORKDIR $GOPATH/src/github.com/sinfo/go-tutorial/
COPY . .

RUN make clean
RUN make deps
RUN make build

RUN apt-get install -y mongodb
CMD ["make", "run"]
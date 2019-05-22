FROM golang:latest

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p $GOPATH/src $GOPATH/bin && chmod -R 777 $GOPATH
WORKDIR $GOPATH

RUN apt-get update
RUN apt-get install git gcc

RUN mkdir -p $GOPATH/src/github/sinfo/hero-tutorial-go
WORKDIR $GOPATH/src/github.com/sinfo/hero-tutorial-go/
COPY . .

RUN make docker-build

CMD ["make", "docker-run"]
FROM golang:1.12.0-alpine3.9

RUN mkdir /go-tutorial
ADD . /go-tutorial
WORKDIR /go-tutorial
RUN go build -o main .
CMD ["/go-tutorial/main"]
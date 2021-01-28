
FROM golang:1.15-alpine AS builder

ADD . /go/src/crawler

WORKDIR /go/src/crawler

RUN go build -mod=vendor -o crawler .

ENTRYPOINT [ "./crawler"]
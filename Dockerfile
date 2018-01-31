FROM golang:1.9.3-alpine

RUN apk --update --no-cache add curl bash git alpine-sdk

WORKDIR /go/src

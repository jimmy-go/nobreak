FROM golang:1.9.3-alpine

RUN apk --update --no-cache add curl bash git alpine-sdk util-linux gcc musl-dev

RUN curl https://glide.sh/get | sh

WORKDIR /go/src

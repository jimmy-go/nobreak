#!/bin/bash

set -o errexit
set -o nounset

BIN=bin/nobreak_amd64

CC=/usr/bin/x86_64-alpine-linux-musl-gcc-6.3.0 CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a \
    --ldflags '-linkmode external -extldflags "-static"' \
    -o $BIN ./cmd/...

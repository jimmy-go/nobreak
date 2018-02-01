#!/bin/bash
set -o errexit
set -o nounset

BIN=bin/nobreak_amd64

ls -lha /usr/local/musl

CC=/usr/local/musl/bin/musl-gcc CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a \
    --ldflags '-linkmode external -extldflags "-static"' \
    -o $BIN ./cmd/...

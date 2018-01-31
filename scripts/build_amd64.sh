#!/bin/bash
set -o errexit
set -o nounset

BIN=bin/nobreak_amd64

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a \
    -o $BIN ./cmd/...

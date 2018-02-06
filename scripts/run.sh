#!/bin/bash
set -o errexit
set -o nounset

BIN=bin/nobreak_temp

CGO_ENABLED=1 go build -a \
    -o $BIN ./cmd/...

$BIN -config=$PWD/_examples/youtube.yml

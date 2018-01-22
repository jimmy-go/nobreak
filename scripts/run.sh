#!/bin/bash
set -o errexit
set -o nounset

BIN=bin/nobreak_temp

CGO_ENABLED=1 go build -a \
    --tags "libsqlite3 darwin" \
    -o $BIN ./cmd/...

$BIN -config=$1

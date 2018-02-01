#!/bin/bash
set -o errexit
set -o nounset

VERSION=$(git describe --always)
PKG="github.com/jimmy-go/nobreak"
IMAGE="nobreak:$VERSION"

docker build -t $IMAGE -f Dockerfile .

docker run \
        --rm -t -p 7070:7070 \
        -v $PWD:/go/src/$PKG \
        -w /go/src/$PKG \
        $IMAGE \
        /bin/bash -c $1

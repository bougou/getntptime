#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

function build() {
  GOOS=$1
  GOARCH=$2
  go build -v -x -o ${SCRIPT_DIR}/getntptime-${GOOS}-${GOARCH}
}

build linux amd64
build linux arm64
build darwin amd64
build darwin arm64

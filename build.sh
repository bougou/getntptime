#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

function build() {
  local os=$1
  local arch=$2
  GOOS=${os} GOARCH=${arch} go build -v -x -o ${SCRIPT_DIR}/getntptime-${os}-${arch}
}

build linux amd64
build linux arm64
build darwin amd64
build darwin arm64

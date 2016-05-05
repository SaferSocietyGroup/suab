#! /bin/bash

set -e

BUILD_CMD="go build -v suab"

READPATH=`readlink -nf $0`
WD=`dirname $READPATH`
export GOPATH=$WD
mkdir -p "${WD}/build"

function build {
  echo "## Compling for $1..."
  GOOS=$1 $BUILD_CMD
  mv -f suab* "${WD}/build"
  echo "## Done"
}


build linux
echo
build windows

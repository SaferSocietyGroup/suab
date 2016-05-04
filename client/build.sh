#! /bin/bash

set -e

mkdir -p build
BUILD_CMD="go build -v suab"

DIRNAME=`dirname $0`
WD=`readlink -nf $DIRNAME`
export GOPATH=$WD

function build {
  echo "## Compling for $1..."
  GOOS=$1 $BUILD_CMD
  mv -f suab* build/
  echo "## Done"
}


build linux
echo
build windows

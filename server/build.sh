#! /bin/bash

set -e

BUILD_CMD="go build -v suab-server"

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

# Generate assets
PATH=$PATH:"./bin"
go-bindata-assetfs ../web-ui/... ../client/build/...
mv bindata_assetfs.go src/suab-server/assetfs.go

build linux
echo
build windows


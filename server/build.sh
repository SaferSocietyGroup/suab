#! /bin/bash

set -e

BUILD_CMD="go build -v suab-server"

READPATH=`readlink -nf $0`
WD=`dirname $READPATH`
export GO15VENDOREXPERIMENT=1
export GOPATH=${WD}
mkdir -p "${WD}/build"

function build {
  echo "## Compling for $1..."
  GOOS=$1 $BUILD_CMD
  mv -f suab* "${WD}/build"
  echo "## Done"
}

# Generate assets
PATH=$PATH:"${WD}/bin"
pushd $WD > /dev/null
test -d ../client/build || (echo "The client wasn't built, building it" ; ../client/build.sh)
go-bindata-assetfs -ignore .*node_modules.* ../web-ui/... ../client/build/...
mv bindata_assetfs.go src/suab-server/assetfs.go
popd > /dev/null

build linux
echo
build windows


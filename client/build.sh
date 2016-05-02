#! /bin/bash

set -e

mkdir -p build
BUILD_CMD="go build -v src/suab.go"

echo "## Compling for linux..."
GOOS=linux $BUILD_CMD
mv -f suab build/
echo "## Done"

echo
echo "## Compiling for Windows"
GOOS=windows $BUILD_CMD
mv -f suab.exe build/
echo "## Done"

#!/usr/bin/env bash

mkdir artifacts
cd cake
./build.sh -t "Zip-Files"

find /cake/artifacts -iname "*.zip" -type f -exec cp {} ../artifacts \\;


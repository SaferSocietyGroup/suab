#!/usr/bin/env sh

## Build the source code
go build -o the-artifact hello

## and move it to /artifacts so that it will be uploaded
mv the-artifact /artifacts

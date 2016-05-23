#!/usr/bin/env bash

mkdir artifacts

cd maven-hello-world/my-app
mvn package
cp target/*.jar ../../artifacts


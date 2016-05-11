#! /bin/bash

WD=`dirname $0`
docker build --rm --tag regression-test-base --file $WD/test-base.dockerfile $WD

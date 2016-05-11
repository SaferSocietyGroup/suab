#! /bin/bash

WD="`dirname $0`"

## Build the base image
$WD/regression/build-regression-test-base.sh > /dev/null || exit 1

EXIT_CODE=0
find $WD/regression -mindepth 1 -maxdepth 1 -type d -print0 | while IFS= read -r -d $'\0' test_dir; do
    $WD/regression/run-test.sh $test_dir
    if [ $? -ne 0 ]; then
      EXIT_CODE=2
    fi

    ## Print a newline between each test
    echo
done
exit $EXIT_CODE

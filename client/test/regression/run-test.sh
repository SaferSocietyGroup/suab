#! /bin/bash
set -u

TEST_DIR=`readlink -f "$1"`
TEST_NAME=`basename $TEST_DIR`
IMAGE_TAG="apa"

SUAB_BINARY="`dirname $0`/../../build/suab"
TEST_CMD="$SUAB_BINARY -d $IMAGE_TAG"

echo "Running test $TEST_NAME"

## RUN TEST
docker build --tag $IMAGE_TAG $TEST_DIR  > /dev/null || exit 1
$TEST_CMD

## CHECK OUTPUT
TEST_EXIT_CODE=$?
EXPECTED_EXIT_CODE=0
if [[ -e $TEST_DIR/expected-exit-code ]]; then
  EXPECTED_EXIT_CODE=`cat $TEST_DIR/expected-exit-code`
fi

#docker rmi $IMAGE_TAG > /dev/null || true

if [ $TEST_EXIT_CODE -eq $EXPECTED_EXIT_CODE ]; then
  echo "TEST $TEST_NAME SUCCESSFUL!"
else
  echo -n "TEST $TEST_NAME FAILED!"

  if [ $EXPECTED_EXIT_CODE -ne 0 ]; then
    echo -n " Expected exit code $EXPECTED_EXIT_CODE but got $TEST_EXIT_CODE "
  fi
  echo

  exit 1
fi


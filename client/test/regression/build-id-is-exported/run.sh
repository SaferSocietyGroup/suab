#! /bin/bash

if [ -n "$SUAB_BUILD_ID" ]; then
  echo "Build id set to $SUAB_BUILD_ID!"
else
  echo "SUAB_BUILD_ID was not set!"
  echo "Available environment variables:"
  env

  exit 1
fi

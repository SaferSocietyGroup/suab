#!/usr/bin/env bash

## This script will checkout the git repository at $REPO_URL and put it in
## $CHECKOUT_DESTINATION. Note that $GIT_COMMIT is expected to passed in as
## an environment variable. This is achieved by the --env flag to suab. If
## omitted, it will checkout master by default

REPO_URL="https://github.com/SaferSocietyGroup/suab.git"
DEFAULT_REVISION="master"
CHECKOUT_DESTINATION="."

set -x
git clone \
	--depth=1 \
	--branch=${GIT_COMMIT:-$DEFAULT_REVISION} \
	$REPO_URL \
	$CHECKOUT_DESTINATION

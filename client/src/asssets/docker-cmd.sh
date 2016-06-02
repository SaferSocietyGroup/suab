#!/usr/bin/env sh

LOG_FILE=/tmp/run-output
EXIT_CODE_FILE=/tmp/suab-exit-code
# Start a subshell and redirect all output to a file
(
    export SUAB_BUILD_ID=`hostname`
    export BASE_URL=http://$1/build/${SUAB_BUILD_ID}
    export IMAGE_TAG=$2
    echo "BuildId: $SUAB_BUILD_ID"

    # TODO: Retries and timeouts...
    curl --data "{\"image\": \"${IMAGE_TAG}\"}" --silent --show-error ${BASE_URL}

    checkout-code.sh
    export CHECKOUT_CODE_EXIT_CODE=$?
    echo ${CHECKOUT_CODE_EXIT_CODE} > ${EXIT_CODE_FILE}
    # TODO: We probably should not let checkout-code.sh and run.sh run forever...
    test ${CHECKOUT_CODE_EXIT_CODE} -eq 0 && (run.sh ; echo $? > ${EXIT_CODE_FILE})

    # Upload logs, TODO: retries and timeout
    curl --data @${LOG_FILE} --silent --show-error ${BASE_URL}/logs

    # Upload artifacts, TODO: retries and timeout
    test -d /artifacts && find /artifacts -type f -exec curl -X POST --data-binary @{} ${BASE_URL}{} \;

    # Upload the exit code, TODO: retries and timeout
    curl --data "{\"exitCode\": \"`cat ${EXIT_CODE_FILE}`\"}" --request PATCH --silent --show-error ${BASE_URL}

) 2>&1 | tee ${LOG_FILE}

# Make sure we exit with the exit code of checkout-code.sh or run.sh
exit `cat ${EXIT_CODE_FILE}`

FROM regression-test-base

ADD run.sh /bin/run.sh
ADD expected-output /tmp/expected-output


CMD ["cmp", "/tmp/run-output", "/tmp/expected-output"]

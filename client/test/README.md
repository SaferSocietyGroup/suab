# Running the tests

## Unit tests
Just run
```
client/test/run-unit-tests.sh
```

## Regression tests

To run them:

1. Start the suab server at `localhost:2397`
2. run `client/test/run-regression-tests.sh`


To add a new one:

1. Create a folder in `client/tests/regression` and name it what you want the test to be called
2. Create a Dockerfile in said folder that exits with 0 if all seems well and non-zero otherwise. It's up to the test to provide enough output to be debuggable. There's a `regression-test-base` docker image that you can use as the `FROM` image.

If you expect an exit code other than zero, put an `expected-exit-code` file with the expected exit code next to the Dockerfile and the test runner will pick it up

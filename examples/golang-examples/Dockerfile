## Start from somewhere convenient for you
FROM golang:1.5

## suab will upload all files in /artifacts if it exists,
## so this is where we'll put the built binary
RUN mkdir /artifacts

## Add the files that will checkout the code and run the build
## They need to be named checkout-code.sh and run.sh and be
## on the path
ADD checkout-code.sh /bin/checkout-code.sh
ADD run.sh /bin/run.sh

## Not needed for suab, but will make it easier for you to
## test your image
CMD run.sh

# suab - The build system you already know

suab is a continuous integration system that tries to get out of the way. You create a docker image that can do whatever it is you do in your builds and suab handles the logs and artifacts for you. It solves the same problem as e.g. Jenkins and Bamboo, but its DSL is bash and everything is under source control.

> **If you can build with bash, you can build in docker. If you can build in docker, you can build it on suab.**

## Usage example
Say we want to build [a simple golang program](https://github.com/eriklarko/golang-hello-world). The directory with all files used looks like this and can be [downloaded here](https://github.com/SaferSocietyGroup/suab/blob/master/examples/golang-hello).

```
golang-hello
├── Dockerfile
├── checkout-code.sh
└── run.sh
```

There are [many more examples here](https://github.com/SaferSocietyGroup/suab/blob/master/examples).

### Dockerfile
suab runs a docker image in which you run your build. Just create the image like you want to.

```dockerfile
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
```
suab only requires `curl`, `find`, `tee`, `echo`, `test`, `exit` and `export`.

### checkout-code.sh
suab runs a script named `checkout-code.sh` to checkout the code you want to build. This script is defined by you and can be as simple as
```bash
#!/usr/bin/env sh
git clone --depth=1 https://github.com/eriklarko/golang-hello-world.git /go/src
```
An example script that can checkout specific revisions can be [found here](https://github.com/SaferSocietyGroup/suab/blob/master/client/clone.sh).

### run.sh
With the source code checked out we can build it. Just put the commands you use to build in a script named `run.sh` and suab will run this for you. 
```bash
#!/usr/bin/env sh

## Build the source code
go build -o the-artifact hello

## and move it to /artifacts so that it will be uploaded
mv the-artifact /artifacts
```
If you want the output from your build to be handled by suab, use `run.sh` to put them in `/artifacts`. Anything outputted to stdout and stderr will automatically be uploaded to the suab server.

### Putting it all together
If you run suab locally, all you need to do now is build the docker image and tell suab to run it.
```bash
$ docker build --tag=suab-example .
$ suab -d suab-example
```
**That's it!**

If you use a separate docker daemon, such as in a docker-swarm, you need to push your image to a repository that can be accessed by the swarm.

### Requirements on the docker image
1. a `checkout-code.sh` script on the path. This is where you get your code into the docker image. An example of this script can be found [here](https://github.com/SaferSocietyGroup/suab/blob/master/client/clone.sh).
2. a `run.sh` script on the path that actually builds your stuff.
3. all logs you are interested in are printed to stdout or stderr
4. any artifacts you want from your build is put in `/artifacts`
5. `curl`, `find`, `tee`, `echo`, `test`, `exit` and `export` on the path.


## Installing the server
1. Have a docker daemon running somewhere. We suggest using a swarm, but your local daemon works fine!
2. Run `docker run -d -p 8080:8080 suab/server`

---

## Building from source
Just run `client/build.sh` or `server/build.sh` and linux and windows binaries will be built and put in the `build` folder of the corresponding projects.

You can also build in docker by running
```
docker build --tag=suab-client-build client
docker run --rm -v `pwd`:/src suab-client-build client/build.sh
```
and
```
docker build --tag=suab-server-build -f server/build.dockerfile server
docker run --rm -v `pwd`:/src suab-server-build server/build.sh
```
from the git root folder. The compiled binaries are now in `client/build` and `server/build`.

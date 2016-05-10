# suab
A ubiquitous-fiesta for your builds. No more *"it works on my machine but breaks on the builder!"*.

No more *"I have to wait for these gazillion other builds to finish before I can build..."*.

**No more having to learn a new build system** - if you can build with bash, you can build in docker. If you can build in docker, you can build it on suab.

## Usage example
1. Create a docker image in which you can build your stuff.
2. run `./suab` with the tag of said docker image and the address to the suab server

**That's it!** 

A gif of this in action is coming up!

No, but really it's easy. Suab expects six things from your docker image

1. `curl` is on the path
2. `find` is on the path
3. a `checkout-code.sh` script on the path. This is where you get your code into the docker image. An example of this script can be found [here](https://github.com/SaferSocietyGroup/suab/blob/master/client/clone.sh).
4. a `run.sh` script on the path that actually builds your stuff.
5. all logs are printed to stdout or stderr
6. any artifacts you want from your build is put in `/artifacts`

There are many examples [here](https://github.com/SaferSocietyGroup/suab/blob/master/examples).

## Installing the server
1. install [docker-compose](https://docs.docker.com/compose/install)
2. download [this docker-compose file](https://github.com/SaferSocietyGroup/suab/blob/master/server/server-compose.yml)
3. run `docker-compose up -d -f THE-FILE-FROM-STEP-2`

## Building from source
Just run `client/build.sh` or `server/build.sh` and linux and windows binary will be built and put in the `build` folder of the corresponding projects.

You can also build in docker by running
```
docker build --tag=suab-client-build client
docker run --rm -v `pwd`:/src suab-client-build client/build.sh
```
and
```
docker build --tag=suab-server-build server
docker run --rm -v `pwd`:/src suab-server-build server/build.sh
```
from the git root folder. The compiled binaries are now in `client/build` and `server/build`.

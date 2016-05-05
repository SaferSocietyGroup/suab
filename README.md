# suab
 ubiquitous-fiesta.


## Compiling the client
Just run
```
$ ./build.sh
```
from the client folder and a linux and windows binary will be built and put in the `build` folder.

You can also build it in docker by running
```
docker build --tag=suab-client-build .
docker run --rm -v `pwd`:/src suab-client-build
```
from the client folder

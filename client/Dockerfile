FROM golang:1.5

COPY clone.sh /bin/checkout-code.sh
COPY build.sh /bin/run.sh

RUN mkdir /artifacts
RUN echo "mv build/* /artifacts" >> /bin/run.sh

ENV GOPATH /src
WORKDIR /src
CMD run.sh

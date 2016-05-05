FROM golang:1.5

RUN mkdir /artifacts
CMD run.sh
ENV GOPATH /src
WORKDIR /src

COPY clone.sh /bin/checkout-code.sh

RUN echo "client/build.sh" >> /bin/run.sh
RUN echo "mv client/build/* /artifacts" >> /bin/run.sh
RUN chmod +x /bin/run.sh

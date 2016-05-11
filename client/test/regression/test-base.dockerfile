FROM ubuntu:16.04

RUN touch /bin/checkout-code.sh && \
    chmod +x /bin/checkout-code.sh

RUN touch /bin/run.sh && \
    chmod +x /bin/run.sh

RUN apt-get update && \
    apt-get install -y curl

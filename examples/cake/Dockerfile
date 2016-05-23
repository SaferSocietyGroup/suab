FROM mono:4.2.3.4

RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

COPY checkout-code.sh /bin/checkout-code.sh
COPY run.sh /bin/run.sh

CMD run.sh
 

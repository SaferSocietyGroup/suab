FROM ubuntu:16.04

ADD build/suab-server /bin/suab-server

CMD suab-server

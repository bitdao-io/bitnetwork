FROM ubuntu:20.04 AS compile-env
FROM golang
ENV GOPROXY=https://goproxy.cn
RUN apt-get install git
RUN cd /root/ && git clone https://github.com/ethandavionlabs/bitnetwork.git && cd bitnetwork && make install
RUN bitnetworkd version


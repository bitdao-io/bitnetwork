FROM golang
ENV GOPROXY=https://goproxy.cn
RUN apt-get install git
RUN cd /root/ && git clone https://github.com/bitdao-io/bitnetwork.git && cd bitnetwork && make install
RUN bitnetworkd version
RUN cd /root/ && git clone https://github.com/bitdao-io/bitos.git && cd bitos && make install


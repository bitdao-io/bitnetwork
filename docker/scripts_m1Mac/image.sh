#!/bin/bash


#operation in compile image
cd /share_file/bitnetwork
make install
mv $GOPATH/bin/bitnetworkd /share_file/


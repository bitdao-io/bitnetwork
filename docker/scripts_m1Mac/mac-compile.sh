#!/bin/bash

IMAGE_NAME=bitchain-compile-image


cd ~/share_file
rm -rf bitnetwork

cp -r ~/bitchain/ethan/bitnetwork ~/share_file
docker exec -it bitchain-generator /share_file/bitnetwork/docker/scripts_m1Mac/image.sh
mv ~/share_file/bitnetworkd $GOPATH/bin











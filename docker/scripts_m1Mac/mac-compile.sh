#!/bin/bash


#path init
rm -rf ~/share_file
mkdir ~/share_file


cd ~/share_file
git clone git@github.com:ethandavionlabs/bitnetwork.git
ls
cd bitnetwork

docker rm -f bitchain-compile
docker run -itd --platform linux/amd64 --name bitchain-compile -v ~/share_file:/workspace -i 8bf155606348 /bin/bash

pwd
#operation in compile image
docker exec -it bitchain-compile ./docker/scripts_m1Mac/image.sh


# outside image
mv ~/share_file/bin/bitnetworkd $GOPATH/bin








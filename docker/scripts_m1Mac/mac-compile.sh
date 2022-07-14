#!/bin/bash


IMAGE_ID=$(docker ps |grep bitchain-compile| awk '{print $1}')


rm -rf ~/share_file
mkdir ~/share_file



if test -z "$IMAGE_ID";then
  echo "start init the docker image"
  docker run -itd  --name=bitchain-compile -v ~/share_file:/share_file wosa/bitchain-compile	 /bin/bash
  echo "finish init the docker image"
else
  docker restart -itd  bitchain-compile
  echo "docker compile image already exists"
fi

rm -rf ~/share_file/bitnetwork
cp -r ../../../bitnetwork ~/share_file/bitnetwork


chmod 777  ~/share_file/bitnetwork/docker/scripts_m1Mac/image.sh
docker exec -it bitchain-compile   /share_file/bitnetwork/docker/scripts_m1Mac/image.sh












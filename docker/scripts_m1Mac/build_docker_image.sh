#!/bin/bash

IMAGE_NAME=bitchain-compile-image


mkdir ~/share_file

docker build -t IMAGE_NAME .

docker run -it IMAGE_NAME --name=bitchain-generator -v ~/share_file:/share_file  /bin/bash

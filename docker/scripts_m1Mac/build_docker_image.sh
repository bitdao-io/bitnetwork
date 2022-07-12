#!/bin/bash

IMAGE_NAME=bitchain-compile-image


mkdir ~/share_file

docker build -t bitchain-compile-image .

docker run -it  --name=bitchain-generator -v ~/share_file:/share_file bitchain-compile-image /bin/bash

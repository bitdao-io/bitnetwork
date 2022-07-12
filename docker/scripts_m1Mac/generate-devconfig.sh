#!/bin/bash





rm -rf validators
mkdir -p validators

# generate new config directory
bitnetworkd testnet init-files --output-dir ../../validators --chain-id bitnetwork_7001-1 \
        --node-daemon-home bitnetwork --keyring-backend test --node-dir-prefix validator \
        --v 7 --starting-ip-address 172.172.0.2

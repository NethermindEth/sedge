#!/bin/bash

# Install docker

sudo apt-get update
sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
sudo apt install -y docker.io

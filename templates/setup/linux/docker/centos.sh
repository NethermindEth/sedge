#!/bin/bash

# Install docker

sudo yum install -y yum-utils
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install -y docker-ce docker-ce-cli containerd.io

# Install docker-compose

DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
mkdir -p $DOCKER_CONFIG/cli-plugins
wget "https://github.com/docker/compose/releases/download/v2.6.1/docker-compose-$(uname -s)-$(uname -m)" 
sudo mv docker-compose-$(uname -s)-$(uname -m) $DOCKER_CONFIG/cli-plugins/docker-compose
sudo chmod -v +x $DOCKER_CONFIG/cli-plugins/docker-compose
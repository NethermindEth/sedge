#!/bin/bash

# Install docker

ARCH=$1

sudo apt-get update
sudo apt install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=$ARCH] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
sudo apt update
sudo apt install -y docker-ce

# Install docker-compose

DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
mkdir -p $DOCKER_CONFIG/cli-plugins
wget "https://github.com/docker/compose/releases/download/v2.6.1/docker-compose-$(uname -s)-$(uname -m)" 
sudo mv docker-compose-$(uname -s)-$(uname -m) $DOCKER_CONFIG/cli-plugins/docker-compose
sudo chmod -v +x $DOCKER_CONFIG/cli-plugins/docker-compose
#!/bin/bash

# Install docker

sudo yum -y update
sudo yum -y install docker

# Start docker engine
sudo service docker start

# Make docker auto-start
sudo chkconfig docker on

# Install docker-compose

DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
mkdir -p $DOCKER_CONFIG/cli-plugins
wget "https://github.com/docker/compose/releases/download/v2.6.1/docker-compose-$(uname -s)-$(uname -m)" 
sudo mv docker-compose-$(uname -s)-$(uname -m) $DOCKER_CONFIG/cli-plugins/docker-compose
sudo chmod -v +x $DOCKER_CONFIG/cli-plugins/docker-compose
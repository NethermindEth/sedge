#!/bin/bash

# Download docker-compose v2.4.1 linux-x86-64 standalone
curl -SL https://github.com/docker/compose/releases/download/v2.4.1/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose

# add execute permission to /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

#test docker-compose 
docker-compose --version
#!/bin/bash

# Install docker

sudo yum -y update
sudo yum -y install docker

# Start docker engine
sudo service docker start

# Make docker auto-start
sudo chkconfig docker on
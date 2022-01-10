#!/bin/bash

# NOTE: Docker is only supported for Fedora 34 and Fedora 35

# Install docker

sudo dnf -y install dnf-plugins-core
sudo dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo
sudo dnf install docker-ce docker-ce-cli containerd.io
#!/bin/sh

# Install Python 3
apk add --update libffi-dev
apk add python3-dev
apk add build-base
apk add --no-cache python3 && ln -sf python3 /usr/bin/python
python3 -m ensurepip
pip3 install --no-cache --upgrade pip setuptools

# Install Docker-Compose
pip3 install docker-compose
#!/bin/sh

# Add Community APK Repo
echo "http://dl-cdn.alpinelinux.org/alpine/v3.15/community" >> /etc/apk/repositories

# Add Docker and Update
apk add docker
apk update

# Run Docker at Boot
rc-update add docker boot

# Start Docker Manually
service docker start
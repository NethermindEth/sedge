#!/bin/bash
sudo pacman -S --noconfirm docker docker-compose
sudo systemctl start docker
sudo systemctl enable docker
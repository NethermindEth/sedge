#!/bin/bash
sudo pacman -S --noconfirm docker
sudo systemctl start docker
sudo systemctl enable docker
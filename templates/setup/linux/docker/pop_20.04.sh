#!/bin/bash


# download the docker bianry package using wget 
wget -c https://download.docker.com/linux/static/stable/x86_64/docker-20.10.14.tgz

# decompress the downloded file 
tar xzvf docker-20.10.14.tgz
rm docker-20.10.14.tgz

#give permission to run docker
sudo chown "$USER":"$USER" /home/"$USER"/.docker -R
sudo chmod g+rwx "$HOME/.docker" -R

# Copy all binaries to /usr/bin
sudo cp docker/* /usr/bin/

# Create Docker group 
sudo groupadd docker

# Add USER to docker's group
sudo usermod -aG docker $USER
#Logout and log back in which your group membership is re-evaluated
sudo newgrp docker 



#Configure Docker to start on boot 

sudo systemctl enable docker.service
sudo systemctl enable containerd.service

rm -rf /home/$USER/docker/
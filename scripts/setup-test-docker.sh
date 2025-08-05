#!/bin/bash

# Setup Docker for test environment
# This script sets up Docker for running tests that require Docker

set -e

echo "Setting up Docker for test environment..."

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "Docker is not installed. Installing Docker..."
    
    # Update package list
    sudo apt-get update
    
    # Install prerequisites
    sudo apt-get install -y \
        ca-certificates \
        curl \
        gnupg \
        lsb-release
    
    # Add Docker's official GPG key
    sudo mkdir -p /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    
    # Set up the repository
    echo \
        "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    
    # Install Docker
    sudo apt-get update
    sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin
    
    echo "Docker installed successfully"
else
    echo "Docker is already installed"
fi

# Start Docker service
echo "Starting Docker service..."
if command -v systemctl &> /dev/null; then
    sudo systemctl start docker
    sudo systemctl enable docker
else
    # If systemctl is not available (e.g., in containers), start Docker manually
    sudo dockerd &
    sleep 5
fi

# Add current user to docker group
echo "Adding user to docker group..."
sudo usermod -aG docker $USER

# Test Docker
echo "Testing Docker..."
if sudo docker ps &> /dev/null; then
    echo "Docker is working correctly"
else
    echo "Warning: Docker might not be fully functional"
fi

# Create a test container to verify everything works
echo "Creating test container..."
if sudo docker run --rm hello-world &> /dev/null; then
    echo "Docker test container created successfully"
else
    echo "Warning: Could not create test container"
fi

echo "Docker setup complete!"
echo ""
echo "Note: You may need to log out and back in for the docker group changes to take effect."
echo "Alternatively, you can run: newgrp docker"
echo ""
echo "To run tests with Docker, use: sudo go test ./... -v"
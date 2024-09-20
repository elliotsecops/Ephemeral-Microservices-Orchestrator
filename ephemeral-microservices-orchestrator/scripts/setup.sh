#!/bin/bash

# Install Go
echo "Installing Go..."
wget https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Install Docker
echo "Installing Docker..."
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER

# Install Minikube
echo "Installing Minikube..."
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

# Start Minikube
echo "Starting Minikube..."
minikube start

# Install kubectl
echo "Installing kubectl..."
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

echo "Setup complete!"
### Only for Macs, PC users—-get a better computer. :P

## Make sure you have VMX enabled 

sysctl -a | grep machdep.cpu.features | grep VMX

## Install brew 
## Update brew
## Install Docker

brew update && brew cask install docker

## Install Minikube 

curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.15.0/minikube-darwin-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/

### You can use minikube with xhyve or VMWare, but I prefer VirtualBox  ¯\_(ツ)_/¯


## Install kubectl

curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl && chmod +x ./kubectl && sudo mv ./kubectl /usr/local/bin/kubectl



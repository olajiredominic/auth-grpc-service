#!/bin/bash
echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/bash.bashrc
source /etc/bash.bashrc

cd /home/ubuntu
chmod +w ./auth_app
cd ./auth_app
go build -o auth_app ./cmd

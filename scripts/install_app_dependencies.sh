#!/bin/bash
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile

cd /home/ubuntu
chmod +w ./auth_app
cd ./auth_app
go build -o auth_app ./cmd

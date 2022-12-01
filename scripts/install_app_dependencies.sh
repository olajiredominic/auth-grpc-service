#!/bin/bash
cd /home/ubuntu
chmod +w ./auth_app
cd ./auth_app
go build -o auth_app ./cmd

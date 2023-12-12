#!/bin/bash
RUN_NAME=videoHost
mkdir -p output/bin output/conf
cp script/bootstrap.sh output 2>/dev/null
chmod +x output/bootstrap.sh
cp -r conf/* output/conf

# Set GOOS and GOARCH for Linux arm64
# export GOOS=linux
# export GOARCH=arm64


go build -o output/bin/${RUN_NAME}
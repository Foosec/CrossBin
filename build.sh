#!/bin/sh

echo "Building for linux"
env GOOS=linux GOARCH=amd64 go build -o dist/CrossBin.bin
echo "Building for windows"
env GOOS=windows GOARCH=amd64 go build -o dist/CrossBin.exe
echo "Building for MacOS"
env GOOS=darwin GOARCH=amd64 go build -o dist/CrossBin.dmg



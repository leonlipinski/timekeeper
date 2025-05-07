#!/bin/bash

GOOS=darwin GOARCH=arm64 go build -o builds/timekeeper-darwin-arm64
GOOS=linux GOARCH=amd64 go build -o builds/timekeeper-linux-amd64
GOOS=windows GOARCH=amd64 go build -o builds/timekeeper-windows-amd64.exe

chmod +x builds/timekeeper-darwin-arm64
chmod +x builds/timekeeper-linux-amd64

#!/bin/bash
echo "Thank you for installing Go Speedtool."
echo "Installing shall take a moment..."
chmod +x ./main.go
export PATH=$PATH:$PWD
$(which go) run "main.go"; exit $?

#!/bin/bash
# Setup the Go workspace.

echo Setting GOPATH to: `pwd`
export GOPATH=`pwd`
echo Adding bin/ to PATH
export PATH=$PATH:$GOPATH/bin

#!/bin/bash
# Build Lewis and put it in the 'bin/' directory.

echo -n "Building Lewis... "
go build -o bin/lewis src/lewis.go
echo "done."

#!/bin/bash

# OS: UNIX/Linux
# bash script to build the src using go-sdk

# Requirement
# Go sdk 1.16.x

# Steps
# run 'bash build.bash -b' in terminal for Building
# run 'bash build.bash -i' in terminal for Installing using Go
# run 'bash build.bash -d' in terminal for getting all dependencies

if command -v "go" >/dev/null 2>&1 ; then # Check go is installed or not 
    if [ "$1" == "-b" ]; then # -b flag to build ghsh
        echo "Building ghsh using go.."
         go build -o ../../ghsh ../../main.go
        echo "Completed"
    elif [ "$1" == "-i" ]; then # -i flag to install ghsh (using go-sdk)
        echo "Building ghsh using go.."
        go install -o ../../ghsh ../../main.go
        echo "Completed"
    elif [ "$1" == "-d" ]; then # -d flag to get all dependencies
        echo "Getting all dependencies.."
        go get ../../...
        echo "Completed"
    fi
else
    echo "It looks like Go isn't installed" # if go isn't installed
fi
#!/bin/bash

if ! command -v go 
then
  echo "Go is not installed. Please install Go first."
  exit 1
fi

echo export GAC=$(pwd) >> ~/.zshrc

go build -o bin/go-app-cli

go install

if [ ! -d storage ]; then
  mkdir -p storage;
fi

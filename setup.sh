#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Function to create directory if it doesn't exist
create_dir() {
    if [ ! -d "$1" ]; then
        mkdir -p "$1"
        echo "Created directory: $1"
    fi
}

# Main project structure
create_dir cmd
create_dir internal
create_dir pkg
create_dir api
create_dir configs
create_dir scripts
create_dir test
create_dir docs
create_dir build
create_dir deployments

# Subfolders
create_dir internal/app
create_dir internal/models
create_dir internal/repository
create_dir internal/service
create_dir internal/middleware

# Create main.go in cmd
touch cmd/main.go
echo "package main

import (
    \"fmt\"
)

func main() {
    fmt.Println(\"Hello, World!\")
}" > cmd/main.go

echo "Created cmd/main.go"

# Create go.mod if it doesn't exist
if [ ! -f "go.mod" ]; then
    go mod init "$(basename $(pwd))"
    echo "Initialized Go module"
fi

echo "Project structure setup complete!"
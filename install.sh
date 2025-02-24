#!/bin/bash

GO_FILE="main.go"
ALIAS_NAME="go-speedtest"

echo "Installing shall take a moment..."

# Detect Go installation path
GO_BIN_PATH=$(go env GOBIN)
if [[ -z "$GO_BIN_PATH" ]]; then
    # If GOBIN is not set, use GOPATH/bin as fallback
    GO_BIN_PATH=$(go env GOPATH)/bin
fi

# Check if Go binary path was found
if [[ -z "$GO_BIN_PATH" ]]; then
    echo "Could not detect Go installation path. Ensure Go is installed and configured correctly."
    exit 1
fi

# Build the Go file
go build -o "$GO_BIN_PATH/$ALIAS_NAME" "$GO_FILE"
if [[ $? -ne 0 ]]; then
    echo "Failed to build $GO_FILE."
    exit 1
fi

# Add Go bin path to PATH if not already added
if ! echo "$PATH" | grep -q "$GO_BIN_PATH"; then
    echo "export PATH=\$PATH:$GO_BIN_PATH" >> "$HOME/.bashrc"
    export PATH="$PATH:$GO_BIN_PATH"
    echo "Added $GO_BIN_PATH to PATH in .bashrc"
else
    echo "$GO_BIN_PATH is already in PATH"
fi

if command -v "$ALIAS_NAME" &> /dev/null; then
    echo "Command '$ALIAS_NAME' is ready to use. You can run it with: $ALIAS_NAME"
    source ~/.bashrc
else
    echo "Something went wrong. Please check the script output for any issues."
fi

echo "Thank you for installing Go Speedtool."
echo "Please restart your PC for the command to work properly."

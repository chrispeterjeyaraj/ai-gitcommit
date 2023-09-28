#!/usr/bin/env bash

# Install the required Go packages
echo "Installing required Go packages..."
go get -u github.com/go-git/go-git
if [ $? -ne 0 ]; then
  echo "Failed to install Go packages."
  exit 1
fi
echo "Go packages installed."

# Find the script file in the repository
script_path=$(find $(pwd) -name "gitcommit.go" | head -1)

if [ -z "$script_path" ]
then
  echo "Script file not found. Please make sure the script is named 'gitcommit.go' and is located in the repository."
  exit 1
fi

# Build the Go executable
echo "Building Go executable..."
go build -o gitcommit "$script_path"
if [ $? -ne 0 ]; then
  echo "Failed to build Go executable."
  exit 1
fi
echo "Go executable built."

# Copy the executable to /usr/local/bin
echo "Copying executable to /usr/local/bin..."
sudo cp gitcommit /usr/local/bin/
echo "Executable copied."

# Define the alias
echo "Defining alias..."
alias aigitcommit="gitcommit"
echo "Alias defined."

# Check if the alias is already present in the .bashrc file
if grep -q "alias aigitcommit='gitcommit'" ~/.bashrc; then
    echo "Alias already present in .bashrc file. Skipping..."
else
    # Add the alias to the .bashrc file
    echo "Adding alias to .bashrc file..."
    echo "alias aigitcommit='gitcommit'" >> ~/.bashrc
    echo "Alias added to .bashrc file."
    # Reload the .bashrc file
    echo "Reloading .bashrc file..."
    source ~/.bashrc
    echo ".bashrc file reloaded."
fi

echo "Installation complete. You can now use the 'aigitcommit' command."
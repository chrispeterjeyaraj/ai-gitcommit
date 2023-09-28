#!/usr/bin/env bash

# Remove the aicommit alias
sed -i '/alias gitcommit="aigitcommit"/d' ~/.bashrc

# Remove the go package
sudo rm /usr/local/bin/gitcommit

echo "aigitcommit has been uninstalled."
#!/bin/bash

# Remove sudo from configure command and execute it
eval "${CONFIG//sudo/}"
# Download Viam AppImage
curl $APPIMAGE -o viam-server
chmod 755 viam-server 
./viam-server --appimage-extract-and-run -config /etc/viam.json
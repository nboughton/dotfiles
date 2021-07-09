#!/bin/bash
yay --devel -Syu
flatpak update
read -rsp $'Press any key to continue...\n' -n1 key


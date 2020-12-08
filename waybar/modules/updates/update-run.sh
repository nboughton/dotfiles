#!/bin/bash
pgrep checkupdates
pac=$?

pgrep yay
aur=$?

if [[ pac -eq 1 ]] && [[ aur -eq 1 ]]; then 
    exit 0
fi

exit 1
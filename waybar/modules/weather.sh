#!/bin/bash
curl wttr.in/?format="%C,+%t\n" | sed -e 's/,.*:/:/'

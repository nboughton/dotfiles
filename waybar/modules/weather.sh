#!/bin/bash
report=$(curl wttr.in/?format="%C,+%t\n" | sed -e 's/,.*:/:/')
echo $report

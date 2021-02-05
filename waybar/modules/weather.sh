#!/bin/bash
report=$(curl -q wttr.in/?format="%C,+%t" 2>/dev/null)
echo $report

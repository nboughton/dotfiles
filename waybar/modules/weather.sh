#!/bin/bash
report=$(curl -q wttr.in/?format="%C,+%t" 2>/dev/null)
if [[ $report =~ "Unknown location" || $report =~ "Sorry" ]]; then
	report="No data"
fi
echo $report

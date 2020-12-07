#!/bin/bash
updates=$(yay --devel -Qqu)
text=$(echo $updates | wc -l)
alt=$text
tooltip=$updates
percentage=$text
class="no-updates"

if [[ $text -gt 0 ]]; then
	class="updates"
fi

echo "{\"text\": \"$text\", \"alt\": \"$alt\", \"tooltip\": \"$tooltip\", \"class\": \"$class\", \"percentage\": $percentage }"

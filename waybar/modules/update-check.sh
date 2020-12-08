#!/bin/bash
aur_updates=$(yay --devel -Qqmu)
pac_updates=$(checkupdates | awk '{print $1}')
text=$(printf "${aur_updates}${pac_updates}" | wc -l)
alt=$text
tooltip=$(printf "${pac_updates}${aur_updates}" | sed ':a;N;$!ba;s/\n/\\n/g')
percentage=$text
class="no-updates"

if [[ $text -gt 0 ]]; then
	class="updates"
fi

echo "{\"text\": \"$text\", \"alt\": \"$alt\", \"tooltip\": \"$tooltip\", \"class\": \"$class\", \"percentage\": $percentage }"

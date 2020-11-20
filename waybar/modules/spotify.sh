#!/bin/sh

class=$(playerctl metadata --player=spotify --format '{{lc(status)}}')
icon=""

if [[ $class == "playing" ]]; then
	info=$(playerctl metadata --player=spotify --format '{{artist}} {{album}} - {{title}}' | sed 's/"//g')
	if [[ ${#info} > 50 ]]; then
		info=$(echo $info | cut -c1-50)"..."
	fi
	text=$icon" "$info
elif [[ $class == "paused" ]]; then
      text=$icon
elif [[ $class == "stopped" ]]; then
	text=""
fi

echo -e "{\"text\":\""$text"\", \"class\":\""$class"\"}"

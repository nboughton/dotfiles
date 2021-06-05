# Notes:

The current config is designed to source some of its color definitions and config files from ~/.cache/wal/ because I'm lazy. Make sure you've generated a colour scheme with pywal before trying to use these dotfiles as is. The config files that are generated through wal have templates that are committed here too.

Most people probably want to disable the custom/auroch module in the waybar config as it's pretty specific to me for letting me know when AUR packages that I maintain are out of date.

If you want to use the custom/updates (obligatory "I use arch btw") module you'll need Go and Yay installed. You'll also need to do the following:

* ````mkdir -p ~/tmp```` as this is where I output the json data that is polled by the waybar config
* Symlink the updates.timer and updates.service files from waybar/modules/updates/ into ~/.config/systemd/user
* Run ````systemctl --user daemon-reload````
* Run ````systemctl --user enable updates.timer````

## Screenshot

![screenshot](/screenshot.png)

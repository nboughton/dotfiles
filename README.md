# Notes:

* GTK Theme: https://github.com/nboughton/NordSur (use sharp branch for Sway to disable rounded window borders)
* waybar taskbar icon theme: https://github.com/nboughton/DarK-fork
* wallpaper: https://i.imgur.com/3RCodwU.jpg

Most people probably want to disable the custom/aurcheck module in the waybar config as it's pretty specific to me for letting me know when AUR NPM packages that I maintain are out of date.

If you want to use the custom/pacman module you'll need Go and Yay installed. You'll also need to do the following:

* ````mkdir -p ~/tmp```` as this is where I output the json data that is polled by the waybar config
* Symlink the updates.timer and updates.service files from waybar/modules/updates/ into ~/.config/systemd/user
* Fix the paths in the .service file for your home directory
* Run ````systemctl --user daemon-reload````
* Run ````systemctl --user enable updates.timer````

## Screenshot

![screenshot](/screenshot.png)

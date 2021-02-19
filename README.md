# Notes:

* GTK Theme: Nordic-darker
* waybar taskbar icon theme: https://github.com/nboughton/DarK-fork (use grep+sed to replace colour ids in the SVG src and then rebuild the SVG icons to get colours)
* wallpaper: https://github.com/dxnst/nord-backgrounds
* Font: Kanit

Most people probably want to disable the custom/auroch module in the waybar config as it's pretty specific to me for letting me know when AUR packages that I maintain are out of date.

If you want to use the custom/updates (obligatory "I use arch btw") module you'll need Go and Yay installed. You'll also need to do the following:

* ````mkdir -p ~/tmp```` as this is where I output the json data that is polled by the waybar config
* Symlink the updates.timer and updates.service files from waybar/modules/updates/ into ~/.config/systemd/user
* Run ````systemctl --user daemon-reload````
* Run ````systemctl --user enable updates.timer````

The reboot/shutdown buttons require Zenity to provide a confirmation dialog.

## Screenshot

![screenshot](/screenshot.png)

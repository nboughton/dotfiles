package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
)

type jsonOutput struct {
	Text       string `json:"text,omitempty"`
	Alt        string `json:"alt,omitempty"`
	Tooltip    string `json:"tooltip,omitempty"`
	Class      string `json:"class,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
}

var outfile = fmt.Sprintf("%s/tmp/updates.json", os.Getenv(("HOME")))

func main() {
	o := jsonOutput{}

	pac, err := exec.Command("checkupdates").CombinedOutput()
	if err != nil && err.Error() == "exit status 1" {
		o.Tooltip = err.Error()
		o.Class = "error"
	}

	aur, err := exec.Command("yay", "--devel", "-Qu").CombinedOutput()
	if err != nil {
		o.Tooltip += err.Error()
		o.Class = "error"
	}

	updates := append(strings.Split(string(pac), "\n"), strings.Split(string(aur), "\n")...)
	updates = removeEmptyLines(updates)

	n := len(updates)

	if o.Class == "error" {
		o.Text = "!"
		o.Alt = "!"
		o.Percentage = 0
	} else {
		o.Class = "no-updates"
		if n > 0 {
			o.Class = "updates"
			o.Tooltip = strings.Join(updates, "\n")

			conn, err := dbus.SessionBusPrivate()
			if err != nil {
				log.Fatal(err)
			}

			if err = conn.Auth(nil); err != nil {
				log.Fatal(err)
			}

			if err = conn.Hello(); err != nil {
				log.Fatal(err)
			}

			// Send notification
			log.Println("Sending notification")
			notify.SendNotification(conn, notify.Notification{
				AppName:       "PACMAN UPDATES",
				ReplacesID:    uint32(0),
				AppIcon:       "mail-message-new",
				Summary:       "Package Updates Available",
				Body:          o.Tooltip,
				Hints:         map[string]dbus.Variant{},
				ExpireTimeout: 10000,
			})

			conn.Close()
		}

		o.Text = fmt.Sprintf("%d", n)
		o.Alt = fmt.Sprintf("%d", n)
		o.Percentage = n
	}

	f, err := os.Create(outfile)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(f).Encode(o)
	f.Close()
}

func removeEmptyLines(list []string) []string {
	out := []string{}
	for _, line := range list {
		if len(line) > 0 {
			out = append(out, line)
		}
	}
	return out
}

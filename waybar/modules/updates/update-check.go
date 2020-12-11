package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/nboughton/dotfiles/waybar/modules/gobar"
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
	m := gobar.Module{
		Name:    "PACMAN UPDATES",
		Summary: "Updates Available",
		JSON:    gobar.JSONOutput{},
	}

	log.Println("Checking Arch repos")
	pac, err := exec.Command("checkupdates").CombinedOutput()
	if err != nil && err.Error() == "exit status 1" {
		m.JSON.Tooltip = err.Error()
		m.JSON.Class = "error"
	}

	log.Println("Checking AUR and devel")
	aur, err := exec.Command("yay", "--devel", "-Qu").CombinedOutput()
	if err != nil {
		m.JSON.Tooltip += err.Error()
		m.JSON.Class = "error"
	}

	updates := append(strings.Split(string(pac), "\n"), strings.Split(string(aur), "\n")...)
	updates = removeEmptyLines(updates)

	n := len(updates)
	m.JSON.Text = fmt.Sprintf("%d", n)
	m.JSON.Alt = fmt.Sprintf("%d", n)
	m.JSON.Percentage = n

	if m.JSON.Class == "error" {
		m.JSON.Text = "!"
		m.JSON.Alt = "!"
		m.JSON.Percentage = 0
	} else {
		m.JSON.Class = "no-updates"
		if n > 0 {
			m.JSON.Class = "updates"
			m.JSON.Tooltip = strings.Join(updates, "\n")
			log.Println("Sending DBUS notification")
			m.Notify(m.JSON.Tooltip, 10000)
		}
	}

	log.Println("Writing JSON file")
	f, err := os.Create(outfile)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	m.JSON.Write(f)
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
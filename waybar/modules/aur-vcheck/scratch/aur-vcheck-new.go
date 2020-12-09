package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/nboughton/dotfiles/waybar/modules/gobar"
	regex "github.com/nboughton/go-utils/regex/common"
)

type config struct {
	Versions []string `json:"versions,omitempty"`
}

type jsonOutput struct {
	Text       string `json:"text,omitempty"`
	Alt        string `json:"alt,omitempty"`
	Tooltip    string `json:"tooltip,omitempty"`
	Class      string `json:"class,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
}

var outfile = fmt.Sprintf("%s/tmp/aur-vcheck.json", os.Getenv(("HOME")))

func main() {
	confPath := "config.json"
	// Check for config path
	if len(os.Args) > 1 {
		confPath = os.Args[1]
	}

	// Read config
	log.Println("Reading config")
	f, err := os.Open(confPath)
	if err != nil {
		log.Fatalf("cannot open config: %s", err)
	}

	c := config{}
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		log.Fatalf("failed to decode config: %s", err)
	}
	f.Close()

	// Check Versions
	log.Println("Checking versions")
	out := []string{}
	for _, appVer := range c.Versions {
		app := fmt.Sprintf("@%s", strings.Split(appVer, "@")[1])
		o, err := exec.Command("npm", "info", app).CombinedOutput()
		if err != nil {
			log.Fatalf("npm exec failed: %s", err)
		}

		v := regex.ANSI.ReplaceAllString(strings.Fields(strings.Split(string(o), "\n")[1])[0], "")
		if appVer != v {
			out = append(out, fmt.Sprintf("%s -> %s", appVer, v))
		}
	}

	n := len(out)
	m := gobar.Module{
		Name:    "AUR VERSION CHECK",
		Summary: "AUR NPM Packages Out Of Date",
		JSON: gobar.JSONOutput{
			Text:       fmt.Sprintf("%d", n),
			Alt:        fmt.Sprintf("%d", n),
			Tooltip:    strings.Join(out, "\n"),
			Class:      "no-updates",
			Percentage: n,
		},
	}

	if n > 0 {
		log.Println("Sending DBUS notification")
		m.JSON.Class = "updates"
		m.Notify(m.JSON.Tooltip, 10000)
	}

	// Write output for waybar module
	log.Println("Writing JSON output")
	f, err = os.Create(outfile)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	m.JSON.Write(f)
}

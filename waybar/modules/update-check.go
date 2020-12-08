package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type jsonOutput struct {
	Text       string `json:"text,omitempty"`
	Alt        string `json:"alt,omitempty"`
	Tooltip    string `json:"tooltip,omitempty"`
	Class      string `json:"class,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
}

func main() {
	pac, err := exec.Command("checkupdates").CombinedOutput()
	if err != nil {
		log.Fatal("could not get pacman updates: ", string(pac))
	}

	aur, err := exec.Command("yay", "--devel", "-Qu").CombinedOutput()
	if err != nil {
		log.Fatal("could not get aur updates: ", string(aur))
	}

	updates := []string{}
	for _, line := range strings.Split(string(pac), "\n") {
		if len(line) > 0 {
			updates = append(updates, line)
		}
	}

	for _, line := range strings.Split(string(aur), "\n") {
		if len(line) > 0 {
			updates = append(updates, line)
		}
	}

	n := len(updates)

	class := "no-updates"
	if n > 0 {
		class = "updates"
	}

	o := jsonOutput{
		Text:       fmt.Sprintf("%d", n),
		Alt:        fmt.Sprintf("%d", n),
		Tooltip:    strings.Join(updates, "\n"),
		Class:      class,
		Percentage: n,
	}

	json.NewEncoder(os.Stdout).Encode(o)
}

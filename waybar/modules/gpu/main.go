package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/nboughton/dotfiles/waybar/modules/gobar"
)

func main() {
	t := gobar.JSONOutput{
		Class: "normal",
	}

	o, _ := exec.Command("sensors", "-A", "amdgpu-pci-0300").CombinedOutput()
	text := strings.Replace(string(o), "+", "", -1)

	t.Text = getTemp(text)

	// Parse the text output into something pretty
	var (
		buf bytes.Buffer
		tw  = tabwriter.NewWriter(&buf, 2, 2, 2, ' ', 0)
	)
	fmt.Fprint(tw, prune(text))
	tw.Flush()

	t.Tooltip = buf.String()
	t.Percentage = getPct(t.Text)

	if t.Percentage >= 90 {
		t.Class = "critical"
	}
	if t.Percentage >= 60 {
		t.Class = "warning"
	}

	t.Write(os.Stdout)
}

func getTemp(s string) string {
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, "edge") {
			return strings.Replace(strings.TrimSpace(strings.Split(line, ":")[1]), ".0", "", 1)
		}
	}

	return ""
}

func getPct(s string) int {
	r := regexp.MustCompile(`(\d+)`)
	i := r.FindAllString(s, 1)[0]
	o, err := strconv.Atoi(i)
	if err != nil {
		return 0
	}
	return o
}

func prune(lines string) string {
	out := ""
	for _, line := range strings.Split(lines, "\n") {
		if line != "" {
			r := regexp.MustCompile(`(\s{2,})`)
			l := strings.TrimSpace(line)
			p := r.ReplaceAllString(l, "\t")
			out += p + "\n"
		}
	}

	return out
}

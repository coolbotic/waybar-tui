package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type WaybarOutput struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Markup  string `json:"markup"`
}

const (
	fgHeader = "#f4b8e4"
)

// fail prints a minimal JSON object so Waybar still shows something.
func fail(msg string) {
	out := WaybarOutput{
		Text: "N/A",
		Tooltip: fmt.Sprintf(
			"<span foreground='%s'>%s</span>",
			fgHeader,
			msg,
		),
		Markup: "pango",
	}
	_ = json.NewEncoder(os.Stdout).Encode(out)
	os.Exit(0)
}

func main() {
	// For now just hard-code something so you can test Waybar integration.
	text := " | ☀️ <span foreground='#a6d189'>20°C</span>"

	tooltip := "" +
		"<span foreground='#f4b8e4' >Current Weather - Demo City</span>\n" +
		"<span foreground='#ffffff'>──────────────</span>\n" +
		"<span foreground='#ffffff' >🌡️ 20°C (Feels like 18°C)</span>\n" +
		"<span foreground='#ffffff'>💧 Humidity: 60%</span>\n"

	out := WaybarOutput{
		Text:    text,
		Tooltip: tooltip,
		Markup:  "pango",
	}

	if err := json.NewEncoder(os.Stdout).Encode(out); err != nil {
		fail(fmt.Sprintf("encode error: %v", err))
	}
}

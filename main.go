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

type WeatherInfo struct {
	Emoji       string
	Description string
}

var WeatherMap = map[int]WeatherInfo{
	0:  {"☀️", "Clear sky"},
	1:  {"🌤️", "Mainly clear"},
	2:  {"⛅", "Partly cloudy"},
	3:  {"☁️", "Overcast"},
	4:  {"☀️", "Clear sky"},
	5:  {"☀️", "Clear sky"},
	45: {"🌫️", "Fog"},
	48: {"🌫️", "Depositing rime fog"},
	51: {"🌦️", "light drizzle"},
	53: {"🌦️", "moderate drizzle"},
	55: {"🌦️", "dense drizzle"},
	61: {"🌧️", "slight rain"},
	63: {"🌧️", "moderate rain"},
	65: {"🌧️", "heavy rain"},
	66: {"🌧️", "light freezing rain"},
	67: {"🌧️", "heavy freezing rain"},
	71: {"❄️", "slight snow"},
	73: {"❄️", "moderate snow"},
	75: {"❄️", "heavy snow"},
	80: {"🌦️", "slight rain showers"},
	81: {"🌧️", "moderate rain showers"},
	82: {"🌧️", "violent rain showers"},
	95: {"⛈️", "Thunderstorm"},
	96: {"⛈️", "Thunderstorm with hail (slight)"},
	99: {"⛈️", "Thunderstorm with hail (severe)"},
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

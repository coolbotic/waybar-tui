package common

import (
	"encoding/json"
	"fmt"
	"os"
)

// Fail prints a minimal JSON object so Waybar still shows something.
func Fail(msg string) {
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

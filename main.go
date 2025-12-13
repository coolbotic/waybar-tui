package main

import (
	"fmt"
	"os"

	"waybar-tui/internal/weather"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: waybar-tui <weather|system>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "weather":
		fmt.Println(weather.Run())
	case "system":
		fmt.Println("James")
	default:
		fmt.Println("Unknow argument...")
		os.Exit(1)
	}
}

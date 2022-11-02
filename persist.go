package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const saveFile = "/var/run/brightness"

func SaveBrightness(b int) {
	payload := []byte(strconv.Itoa(b) + "\n")
	err := os.WriteFile(saveFile, payload, 0644)
	if err != nil {
		log.Printf("Cannot save brightness value to %s: %s", saveFile, err)
	}
}

func LoadBrightness() int {
	contents, err := os.ReadFile(saveFile)
	if err != nil {
		log.Printf("Cannot load brightness value from %s: %s", saveFile, err)
		return 100
	}
	value, err := strconv.Atoi(strings.TrimSpace(string(contents)))
	if err != nil {
		log.Printf("Cannot load brightness value from %s: %s", saveFile, err)
		return 100
	}
	if value > 100 || value < 0 {
		return 100
	}
	return value
}

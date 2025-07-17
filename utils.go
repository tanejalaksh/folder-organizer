package main

import (
	"encoding/json"
	"log"
	"os"
)

func checkFolder(s Settings, ext string) string {
	for f := range s["type"] {
		for i := 0; i < len(s["type"][f]); i++ {
			if ext == s["type"][f][i] {
				return f
			}
		}
	}
	return "undefined"
}

func readSettings(s *Settings) {
	settings, err := os.ReadFile("./settings.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(settings, s); err != nil {
		log.Fatal(err)
	}
}

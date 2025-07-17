package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {

	s := make(map[string]map[string][]string)

	settings, err := os.ReadFile("./settings.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(settings, &s); err != nil {
		log.Fatal(err)
	}

}

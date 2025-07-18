package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type folderLog struct {
	name  string
	count int
}

func checkFolderName(s Settings, ext string) string {
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

func checkFolderLogs(folderLogs []folderLog, folder string) (bool, int) {
	for i, folderLog := range folderLogs {
		if folderLog.name == folder {
			return true, i
		}
	}
	return false, 0
}

func printFolderLogs(folderLog []folderLog) {
	for _, folderCount := range folderLog {
		fmt.Println("\t", folderCount.count, "files", "=>", folderCount.name)
	}
}

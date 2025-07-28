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
type Settings struct {
	Type               map[string][]string `json:"type"`
	SubfolderByMonth   bool                `json:"subfolderByMonth"`
	BypassOrganizeType bool                `json:"bypassOrganizeType"`
}

func initialize(dir string) ([]os.DirEntry, string) {
	dirInfo, err := os.Stat(dir)
	if err != nil || !dirInfo.IsDir() {
		log.Fatal("System couldn't find directory")
	}
	path := dirInfo.Name()

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return files, path

}

func checkFolderName(s map[string][]string, ext string) string {
	for f := range s {
		for i := 0; i < len(s[f]); i++ {
			if ext == s[f][i] {
				return f
			}
		}
	}
	return "/undefined"
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

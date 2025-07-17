package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func organizeByType(rootDir os.FileInfo, files []os.DirEntry, s Settings) {
	for _, entry := range files {
		if entry.IsDir() {
			continue
		}
		folder := checkFolder(s, filepath.Ext(entry.Name()))
		fmt.Println(folder, entry.Name())
	}
}

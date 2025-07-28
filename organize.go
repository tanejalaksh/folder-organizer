package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func organizeByType(path string, files []os.DirEntry, s map[string][]string) []folderLog {
	folderLogs := make([]folderLog, 0, 100)

	for _, entry := range files {
		if entry.IsDir() {
			continue
		}
		folderName := checkFolderName(s, filepath.Ext(entry.Name()))
		dir := path + folderName

		oldPath := path + "/" + entry.Name()
		newPath := dir + "/" + entry.Name()

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal(err)

		}
		if err := os.Rename(oldPath, newPath); err != nil {
			log.Fatal(err)
		}
		if exist, pos := checkFolderLogs(folderLogs, folderName); exist {
			folderLogs[pos].count++
		} else {
			folderLogs = append(folderLogs, folderLog{folderName, 1})
		}
	}
	return folderLogs
}

func organizeByDate(path string, files []os.DirEntry, monthly bool) []folderLog {
	folderLogs := make([]folderLog, 0, 100)
	for _, entry := range files {
		if entry.IsDir() {
			continue
		}
		file, err := entry.Info()
		if err != nil {
			log.Println(err)
			continue
		}

		var folderName string

		if monthly {
			folderName = file.ModTime().Month().String() + " "
		}

		folderName += strconv.Itoa(file.ModTime().Year())

		dir := path + "/" + folderName

		oldPath := path + "/" + entry.Name()
		newPath := dir + "/" + entry.Name()

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal(err)

		}
		if err := os.Rename(oldPath, newPath); err != nil {
			log.Fatal(err)
		}
		if exist, pos := checkFolderLogs(folderLogs, folderName); exist {
			folderLogs[pos].count++
		} else {
			folderLogs = append(folderLogs, folderLog{folderName, 1})
		}
	}

	return folderLogs
}

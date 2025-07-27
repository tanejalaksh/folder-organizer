package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Settings struct {
	Type             map[string][]string `json:"type"`
	SubfolderByMonth bool                `json:"subfolderByMonth"`
}

func main() {
	isTime := flag.Bool("time", false, "organize by date modified")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("Incorrect usage\nCorrect usage:\norganize <dir>\norganize --time <dir>")
	}

	dirInfo, err := os.Stat(flag.Arg(0))
	if err != nil || !dirInfo.IsDir() {
		log.Fatal("System couldn't find directory")
	}

	var s Settings
	readSettings(&s)
	path := dirInfo.Name()

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Organizing", path)

	if len(files) == 0 {
		fmt.Println("Empty dir")
		return
	}

	if *isTime {
		folderLogs := organizeByType(path, files, s.Type)
		printFolderLogs(folderLogs)
		for _, f := range folderLogs {
			dir := dirInfo.Name() + "/" + f.name
			subFiles, err := os.ReadDir(dir)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("organizing", f.name)
			organizeByDate(dir, subFiles, s.SubfolderByMonth)
		}

	} else {
		folderLogs := organizeByType(dirInfo.Name(), files, s.Type)
		printFolderLogs(folderLogs)
	}
}

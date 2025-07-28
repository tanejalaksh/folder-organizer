package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	isTime := flag.Bool("time", false, "organize by date modified")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("Incorrect usage\nCorrect usage:\norganize <dir>\norganize --time <dir>")
	}

	var s Settings
	readSettings(&s)

	files, path := initialize(flag.Arg(0))

	fmt.Println("Organizing", path)

	if len(files) == 0 {
		fmt.Println("Empty dir")
		return
	}

	if *isTime {
		if s.BypassOrganizeType {
			organizeByDate(path, files, s.SubfolderByMonth)
		}
		folderLogs := organizeByType(path, files, s.Type)
		printFolderLogs(folderLogs)
		for _, f := range folderLogs {
			dir := path + "/" + f.name
			subFiles, err := os.ReadDir(dir)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Organizing", f.name)
			folderLogs := organizeByDate(dir, subFiles, s.SubfolderByMonth)
			printFolderLogs(folderLogs)
		}

	} else {
		folderLogs := organizeByType(path, files, s.Type)
		printFolderLogs(folderLogs)
	}
}

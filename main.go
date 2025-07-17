package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Settings map[string]map[string][]string

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

	files, err := os.ReadDir(dirInfo.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Organizing", dirInfo.Name())

	if len(files) == 0 {
		fmt.Println("Done")
		return
	}

	if *isTime {
		return
	} else {
		organizeByType(dirInfo, files, s)
	}

}

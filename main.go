package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Time struct {
	SubfolderByType bool `json:"subfolderByType"`
	Monthly         bool `json:"monthly"`
}

type Settings struct {
	Type map[string][]string `json:"type"`
	Time Time                `json:"time"`
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

	files, err := os.ReadDir(dirInfo.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Organizing", dirInfo.Name())

	if len(files) == 0 {
		fmt.Println("Empty dir")
		return
	}

	if *isTime {
		organizeByDate(dirInfo, files, s.Time)
	} else {
		organizeByType(dirInfo, files, s.Type)
	}

}

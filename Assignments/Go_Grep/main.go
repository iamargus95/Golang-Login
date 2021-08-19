package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var search = ""
var count = 0

func check(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}

func examine(path string, f os.FileInfo, err error) error {
	file, err1 := os.Open(path)
	check(err1)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), search) {
			fmt.Println(path)
			count += 1
		}
	}
	return nil
}

type option struct {
	filename, caseInsensitive, lineNumber, wholeLine bool
}

func getOptions(flags []string) option {
	o := option{}
	for _, f := range flags {
		switch f {
		case "-i":
			o.caseInsensitive = true
		case "-n":
			o.lineNumber = true
		case "-l":
			o.filename = true
		case "-x":
			o.wholeLine = true
		}
	}
	return o
}

func main() {
	fmt.Println()
	if len(os.Args) <= 2 {
		fmt.Println("You must enter both parameters.")
		fmt.Println("Ex. go run main.go <PATH> \"pattern match\"")
		os.Exit(1)
	}

	dir := os.Args[1]
	search = os.Args[2]
	fmt.Println("Results:\n========")
	err := filepath.Walk(dir, examine)
	check(err)
	if count == 0 {
		fmt.Println("No matches found.")
	}
}

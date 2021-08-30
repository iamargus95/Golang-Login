package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {

	var after int
	flag.IntVar(&after, "A", 2, "number of lines after the match.")

	var before int
	flag.IntVar(&before, "B", 2, "number of lines before the match.")

	var around int
	flag.IntVar(&around, "C", 2, "number of lines before and after the match.")

	var recursive bool
	flag.BoolVar(&recursive, "r", true, "List files in a directory recursively.")

	var insensitive bool
	flag.BoolVar(&insensitive, "i", true, "List files in a directory recursively.")

	flag.Parse()

	pattern := patternString()
	path := os.Args[2]
	fmt.Printf("\n" + searchString(path, pattern) + "\n")
}

func walkPathStore() []string {

	inputFiles := []string{}
	err := filepath.Walk(os.Args[2], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			inputFiles = append(inputFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return inputFiles
}

func searchString(path, pattern string) string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(pattern) // this can also be a regex

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			return (scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return (scanner.Text())
}

func patternString() string {
	return (os.Args[1])
}

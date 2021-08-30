package gogrep

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

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

func SearchString(path, pattern string) string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if *insensitive == false {
		r, err := regexp.Compile("(?i)" + pattern) // this can also be a regex
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
	} else {
		r, err := regexp.Compile(pattern)
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
	return (scanner.Text())
}

func PatternString() string {
	return (os.Args[1])
}

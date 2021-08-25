package gogrep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Search(pattern string, flags []string, files []string) []string {
	result := []string{}
	lineNumbers := hasFlag(flags, "-n")
	fileNames := hasFlag(flags, "-l")
	invert := hasFlag(flags, "-v")
	multiFile := len(files) > 1
	exp, e := compileRegexp(pattern, flags)
	if e != nil {
		log.Fatal(e)
		return nil
	}

	for _, file := range files {
		f, e := os.Open(file)
		if e != nil {
			log.Fatal(e)
			return nil
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		lineNo := 0
		for buf, _, e := reader.ReadLine(); e == nil; buf, _, e = reader.ReadLine() {
			line := string(buf)
			lineNo++
			if exp.MatchString(line) != invert {
				if lineNumbers {
					line = fmt.Sprintf("%d:%s", lineNo, line)
				}
				if multiFile && !fileNames {
					line = fmt.Sprintf("%s:%s", file, line)
				}

				if fileNames && !contains(result, file) {
					result = append(result, file)
				} else if !fileNames {
					result = append(result, line)
				}
			}
		}
	}
	return result
}

func hasFlag(flags []string, flag string) bool {
	return contains(flags, flag)
}

func contains(values []string, x string) bool {
	for _, f := range values {
		if f == x {
			return true
		}
	}
	return false
}

func compileRegexp(pattern string, flags []string) (*regexp.Regexp, error) {

	caseInsensitive := hasFlag(flags, "-i")
	entireLines := hasFlag(flags, "-x")
	if entireLines {
		pattern = fmt.Sprintf("^%s$", pattern)
	}
	if caseInsensitive {
		pattern = "(?i)" + pattern
	}

	return regexp.Compile(pattern)
}

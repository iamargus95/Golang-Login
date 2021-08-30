package main

import (
	"flag"
	"fmt"
	"os"
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

	flag.Parse()

	pattern := PatternString()
	path := os.Args[2]
	fmt.Printf(SearchString(path, pattern) + "\n")
}

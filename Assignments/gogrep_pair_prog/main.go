package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func walkpath(path string, f os.FileInfo, err error) error {
	fmt.Printf("%s with %d bytes\n", path, f.Size())
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0) // 1st argument is the directory location
	filepath.Walk(root, walkpath)
}

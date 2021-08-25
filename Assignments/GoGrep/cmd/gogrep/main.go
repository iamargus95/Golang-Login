package main

import (
	"fmt"
	"os"

	"github.com/google/shlex"
	gogrep "github.com/iamargus95/Learning-Golang/Assignments/GoGrep/GoGrep"
)

func main() {
	pattern := os.Args[1]
	flags, _ := shlex.Split(os.Args[2])
	files, _ := shlex.Split(os.Args[3])
	result := gogrep.Search(pattern, flags, files)
	fmt.Println(pattern, flags, files)
	fmt.Println(result)
}

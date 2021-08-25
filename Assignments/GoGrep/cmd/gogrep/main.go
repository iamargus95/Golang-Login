package main

import (
	"fmt"
	"os"

	"github.com/google/shlex"
	gogrep "github.com/iamargus95/Learning-Golang/Assignments/GoGrep/GoGrep"
)

func main() {
	// fmt.Println(os.Args[3], os.Args[4], os.Args[5])
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[2])

	pattern := os.Args[1]
	flags, _ := shlex.Split(os.Args[2])
	files, _ := shlex.Split(os.Args[3])
	result := gogrep.Search(pattern, flags, files)
	fmt.Println(result)
}

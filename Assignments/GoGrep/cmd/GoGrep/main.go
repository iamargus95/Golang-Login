package main

import (
	"fmt"
	"os"

	"github.com/google/shlex"
	"github.com/iamargus95/Learning-Golang/Assignments/GoGrep/GoGrep"
)

func main() {
	flags, _ := shlex.Split(os.Args[2])
	files, _ := shlex.Split(os.Args[3])
	result := GoGrep.Search(os.Args[1], flags, files)
	fmt.Println(result)
}

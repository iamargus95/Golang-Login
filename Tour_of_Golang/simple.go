package main

import "fmt"

func main() {
	fmt.Println(add(2, 3))
	// fmt.Println(add(4, 5))
	// fmt.Println(add(8, 20))
}

func add(a, b int) int {
	defer fmt.Println("The result of the addition is :")
	// return (a + b)
	return (add1(a, b))

}

func add1(a, b int) int {
	fmt.Println("This is add1: ")
	return (a + b)
}

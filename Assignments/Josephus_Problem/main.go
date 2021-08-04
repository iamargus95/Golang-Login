package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	N int //Nunber of people in the Josephus Problem.
)

func main() {

	fmt.Printf("\nEnter the number of people you want for the josephus problem: ")

	fmt.Scanln(&N)
	DecimalN := int64(N)

	BinaryS := strconv.FormatInt(DecimalN, 2)

	BinaryS_array := strings.Split(BinaryS, "")

	first_element, _ := strconv.Atoi(BinaryS_array[0])
	To_be_moved := []int{first_element}

	BinaryN_array := []int{}
	for _, i := range BinaryS_array {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}

		BinaryN_array = append(BinaryN_array, j)
	}

	len := len(BinaryN_array)

	Rest_of_the_array := BinaryN_array[1:len]
	Wb := append(Rest_of_the_array, To_be_moved...)

	Wb_string := []string{}
	for _, i := range Wb {
		j := strconv.Itoa(i)
		Wb_string = append(Wb_string, j)
	}

	Wb_final := strings.Join(Wb_string, "")

	Wb_final_Decimal, _ := strconv.ParseInt(Wb_final, 2, 64)
	fmt.Printf("\n\nThe position of the last surviving person is : %d\n\n", Wb_final_Decimal)
}

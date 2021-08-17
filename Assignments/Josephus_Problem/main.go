package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	N int //Nunber of soldiers in the Josephus Problem.
)

func main() {

	fmt.Printf("\nEnter the number of people you want for the josephus problem: ")

	fmt.Scanln(&N)
	decimalN := int64(N)

	binarySoldier := strconv.FormatInt(decimalN, 2)

	binarySoldierArray := strings.Split(binarySoldier, "")

	firstElement, _ := strconv.Atoi(binarySoldierArray[0])
	toBeMoved := []int{firstElement}

	mutatedBinarySoldierArray := []int{}
	for _, i := range binarySoldierArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}

		mutatedBinarySoldierArray = append(mutatedBinarySoldierArray, j)
	}

	len := len(mutatedBinarySoldierArray)

	restOfTheArray := mutatedBinarySoldierArray[1:len]
	survivorSoldier := append(restOfTheArray, toBeMoved...)

	survivorSoldierString := []string{}
	for _, i := range survivorSoldier {
		j := strconv.Itoa(i)
		survivorSoldierString = append(survivorSoldierString, j)
	}

	survivorFinal := strings.Join(survivorSoldierString, "")

	survivorFinalInt, _ := strconv.ParseInt(survivorFinal, 2, 64)
	fmt.Printf("\n\nThe position of the last surviving person is : %d\n\n", survivorFinalInt)
}

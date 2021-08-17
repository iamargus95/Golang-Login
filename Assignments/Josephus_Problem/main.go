package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var noOfSoldiers int64
	fmt.Printf("\nEnter the number of people you want for the josephus problem: ")

	fmt.Scanln(&noOfSoldiers)

	josephusSimulation(noOfSoldiers)
}

func josephusSimulation(noOfSoldiers int64) {

	binarySoldier := strconv.FormatInt(noOfSoldiers, 2)
	fmt.Println(binarySoldier)

	binarySoldierArray := strings.Split(binarySoldier, "")

	firstElement, _ := strconv.Atoi(binarySoldierArray[0])

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
	survivorSoldier := append(restOfTheArray, firstElement)

	survivorSoldierString := []string{}
	for _, i := range survivorSoldier {
		j := strconv.Itoa(i)
		survivorSoldierString = append(survivorSoldierString, j)
	}

	survivorFinal := strings.Join(survivorSoldierString, "")

	survivorFinalInt, _ := strconv.ParseInt(survivorFinal, 2, 64)
	fmt.Printf("\n\nThe position of the last surviving person is : %d\n\n", survivorFinalInt)
}

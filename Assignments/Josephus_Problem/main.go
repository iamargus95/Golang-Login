package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	collectInput()

}

func collectInput() string {
	var collectStr string
	fmt.Printf("\n Enter the number of people you want for the josephus problem: ")
	fmt.Scanln(&collectStr)

	validateInput(collectStr)

	return collectStr

}

func validateInput(collectStr string) {
	var validInt int
	validInt, err := strconv.Atoi(collectStr)
	if err != nil {
		fmt.Printf("\n Input value should be of type [int],\n ERROR: %q\n", err)
		collectInput()
	} else {
		fmt.Printf("\n %d -> Is expected Input\n", validInt)
		josephusSimulation(validInt)
	}

}

func josephusSimulation(validInt int) int {

	strSoldier := strconv.Itoa(validInt)

	binarySoldier, _ := strconv.ParseInt(strSoldier, 10, 64)

	binarySoldierStr := strconv.FormatInt(binarySoldier, 2)

	binarySoldierArray := strings.Split(binarySoldierStr, "")

	firstElement, _ := strconv.Atoi(binarySoldierArray[0])

	mutatedBinarySoldierArray := []int{}
	for _, i := range binarySoldierArray {
		j, _ := strconv.Atoi(i)
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
	fmt.Printf("\n The position of the last surviving soldier is No: %d\n\n", survivorFinalInt)

	return int(survivorFinalInt)
}

package main

import (
	"fmt"
	"testing"
)

func TestJosephusSimulation(t *testing.T) {
	output := josephusSimulation(41)
	if output == 19 {
		fmt.Printf("Test passed. Expected value returned.")
	}
}

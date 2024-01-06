package main

import (
	"fmt"
	"os"

	diceroller "github.com/skyestalimit/diceroller"
)

func main() {
	// Validate args have been captured
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments: ", len(os.Args)-1)
		printUsage()
		return
	}

	// Captured args sent to be parsed into DiceRolls
	rollInput := os.Args[1:len(os.Args)]
	diceRolls, argErrs := diceroller.ParseRollArgs(rollInput)

	// Print out parsing errors
	for i := range argErrs {
		fmt.Println(argErrs[i])
	}

	// Roll!
	rollsResult, diceErrs := diceroller.PerformRolls(diceRolls)
	if len(diceErrs) > 0 {
		strErrs := ""
		for i := range diceErrs {
			strErrs += diceErrs[i].Error()
		}
		fmt.Println("Unexpected dice roll errors:" + strErrs)
	}

	// Print results
	for i := range rollsResult {
		fmt.Println(rollsResult[i].String())
	}
	fmt.Println("Rolls sum:", diceroller.DiceRollResultsSum(rollsResult))
}

func printUsage() {
	fmt.Println("Usage:	dndhelper [DiceRoll] ... [DiceRoll]")
	fmt.Println("	DiceRoll examples: 1d6, 4d4, 1d10+1, 1D8-1")
	fmt.Println("	Returns the sum of all DiceRolls")
}

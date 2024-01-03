package main

import (
	"fmt"
	"os"

	roller "github.com/diceroller"
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
	argsMap := roller.ParseRollArgs(rollInput)

	// Keep valid DiceRolls
	diceRolls := make([]roller.DiceRoll, 0)
	for diceRoll := range argsMap {
		if argErr := argsMap[diceRoll]; argErr == nil {
			diceRolls = append(diceRolls, *diceRoll)
		} else {
			fmt.Println(argErr)
		}
	}

	// Roll!
	rollResult := roller.PerformRolls(diceRolls)

	// Print results
	fmt.Println("Rolls sum:", getRollsSum(rollResult))
}

func getRollsSum(rollMap map[*roller.DiceRoll]*roller.DiceRollResult) (sum int) {
	sum = 0
	for i := range rollMap {
		result := rollMap[i]
		sum += result.Sum
	}
	return sum
}

func printUsage() {
	fmt.Println("Usage:	dndhelper [DiceRoll] ... [DiceRoll]")
	fmt.Println("	DiceRoll examples: 1d6, 4d4, 1d10+1, 1D8-1")
	fmt.Println("	Returns the sum of all DiceRolls")
}

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
	// capture a dice roll input and roll it
	rollInput := os.Args[1:len(os.Args)]
	rollResult := roller.PerformRolls(roller.ParseRollArgs(rollInput))
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

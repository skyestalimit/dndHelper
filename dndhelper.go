/*
Dndhelper takes command line arguments, sends them to a roll arg parser,
and attempts to perform the rolls of the resulting DiceRolls.

The results of each DiceRoll will be displayed along with the total sum
of every DiceRoll.

Usage:

	dndhelper [rollArg] ... [rollArg]

A rollArg must be in the format XdY(-|+Z). You can send as many rollArg as you want.

Examples of valid roll args:

	1d6, 4d4, 1d10+1, 1D8-1
*/
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
	rollArgs := os.Args[1:len(os.Args)]
	diceRolls, argErrs := diceroller.ParseRollArgs(rollArgs)

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
	fmt.Println("Usage:	dndhelper [rollArg] ... [rollArg]")
	fmt.Println("	Valid rollArg examples: 1d6, 4d4, 1d10+1, 1D8-1")
	fmt.Println("	Returns the sum of all DiceRolls")
}

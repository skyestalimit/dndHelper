package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"

	roller "github.com/diceroller"
)

func main() {
	// capture a dice roll input and roll it
	rollInput := os.Args[1:len(os.Args)]
	fmt.Println("rollInput: ", rollInput)
	rollResult := roller.PerformRolls(parseRollArgs(rollInput))
	fmt.Printf("Rolls result: %v \n", rollResult)
}

func parseRollArgs(rollArgs []string) []roller.DiceRoll {
	// Parse the dice roll and return the ammount and size of dice
	diceRolls := make([]roller.DiceRoll, 0)
	for i := 0; i < len(rollArgs); i++ {
		diceRolls = append(diceRolls, parseRollArg(rollArgs[i]))
	}
	return diceRolls
}

func parseRollArg(rollArg string) roller.DiceRoll {
	// Parse a single roll argument and returns it as a DiceRoll
	argSlices := strings.Split(strings.ToLower(rollArg), "d")
	dice, diceErr := strconv.Atoi(argSlices[0])
	if diceErr != nil {
		fmt.Println(fmt.Sprintf("Error converting dice value of %s: %s", argSlices, diceErr.Error()))
		dice = 0
	}
	ammount, ammountErr := strconv.Atoi(argSlices[1])
	if ammountErr != nil {
		fmt.Println(fmt.Sprintf("Error converting ammount value of %s: %s", argSlices, ammountErr.Error()))
		ammount = 0
	}
	return roller.DiceRoll{dice, ammount, 0}
}
package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"regexp"
	"errors"

	roller "github.com/diceroller"
)

const plus = "+"
const minus = "-"

func main() {
	// Validate args have been captured
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments: ", len(os.Args)-1)
		printUsage()
		return
	}
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
		fmt.Println("Evaluating roll argument: ", rollArgs[i])
		diceRoll, err := parseRollArg(rollArgs[i])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Appending dice roll: ", diceRoll)
			diceRolls = append(diceRolls, diceRoll)
		}
	}
	return diceRolls
}

func parseRollArg(rollArg string) (roller.DiceRoll, error) {
	// Validate arg format
	regExp := regexp.MustCompile("^[[:digit:]]+d[[:digit:]]+([+|-][[:digit:]])?$")
	if !regExp.MatchString(rollArg) {
		return roller.DiceRoll{0, 0, 0}, errors.New(fmt.Sprintf("Invalid roll arg: %s", rollArg))
	}
	// Parse a single roll argument and returns it as a DiceRoll
	rollArg, modifier := evaluateModifier(rollArg)
	
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
	return roller.DiceRoll{dice, ammount, modifier}, nil
}

func evaluateModifier(rollArg string) (string, int) {
	mod := 0
	if strings.ContainsAny(rollArg, plus){
		rollArg, mod = parseModifier(rollArg, plus)
		fmt.Println("Contains +")

	} else if strings.ContainsAny(rollArg, minus) {
		rollArg, mod = parseModifier(rollArg, minus)
		fmt.Println("Contains -")
	}
	return rollArg, mod
}

func parseModifier(rollArg string, symbol string) (string, int){
	modSlices := strings.Split(rollArg, symbol)
	mod, modErr := strconv.Atoi(modSlices[1])
	if modErr != nil {
		fmt.Println("Error converting modifier: ", modSlices[1])
		mod = 0
	} else {
		rollArg = modSlices[0]
		if strings.EqualFold(symbol, minus){
			 mod = -mod
		}
	}
	return rollArg, mod
}

func printUsage() {
	fmt.Println("Usage:	dndhelper XdX ... XdX")
	fmt.Println("	returns the result of the specific rolls")
}
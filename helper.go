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
	rollResult := roller.PerformRolls(parseRollArgs(rollInput))
	fmt.Printf("Rolls result: %v \n", rollResult)
}

func parseRollArgs(rollArgs []string) []roller.DiceRoll {
	// Parse the dice roll and return the ammount and size of dice
	diceRolls := make([]roller.DiceRoll, 0)
	for i := 0; i < len(rollArgs); i++ {
		diceRoll, err := parseRollArg(rollArgs[i])
		if err != nil {
			fmt.Println(err)
		} else {
			diceRolls = append(diceRolls, diceRoll)
		}
	}
	return diceRolls
}

func parseRollArg(rollArg string) (roller.DiceRoll, error) {
	// Validate arg format
	rollArg = strings.ToLower(rollArg)
	regExp := regexp.MustCompile("^[[:digit:]]+d[[:digit:]]+([+|-][[:digit:]])?$")
	if !regExp.MatchString(rollArg) {
		return roller.DiceRoll{0, 0, 0}, errors.New(fmt.Sprintf("Invalid roll arg: %s", rollArg))
	}
	// Parse a single roll argument and returns it as a DiceRoll
	rollArg, modifier := evaluateModifier(rollArg)
	diceAmmount, diceSize := evaluateDiceSizeAndAmmount(rollArg)

	return roller.DiceRoll {diceAmmount, diceSize, modifier}, nil
}

func evaluateModifier(rollArg string) (string, int) {
	mod := 0
	if strings.ContainsAny(rollArg, plus){
		rollArg, mod = parseModifier(rollArg, plus)

	} else if strings.ContainsAny(rollArg, minus) {
		rollArg, mod = parseModifier(rollArg, minus)
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

func evaluateDiceSizeAndAmmount(rollArg string) (int, int){
	argSlices := strings.Split(rollArg, "d")
	ammount := parseDiceSlice(argSlices[0])
	size := parseDiceSlice(argSlices[1])
	return ammount, size
}

func parseDiceSlice(diceSlice string) int {
	dice, diceErr := strconv.Atoi(diceSlice)
	if diceErr != nil {
		fmt.Println(fmt.Sprintf("Error converting dice value of %s: %s", diceSlice, diceErr.Error()))
		dice = 0
	}
	return dice
}

func printUsage() {
	fmt.Println("Usage:	dndhelper XdX ... XdX")
	fmt.Println("	returns the result of the specific rolls")
}
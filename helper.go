package main

import (
	"fmt"

	roller "github.com/dndRoller"
)

func main() {
	// capture a dice roll input and roll it
	rollInput := "1d8"
	rollResult := roller.RollSingleDice(rollInput)
	fmt.Printf("Roll result: %v", rollResult)
}

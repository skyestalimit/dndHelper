# dndhelper

Simple command line app to use the diceroller lib.

## Usage

Dndhelper takes command line arguments, sends them to a roll arg parser,
and attempts to perform the rolls of the resulting DiceRolls.

The results of each DiceRoll will be displayed along with the total sum
of every DiceRoll.

Usage:

dndhelper [rollArg] ... [rollArg]

A rollArg is in the format XdY(-|+Z). You can send as many rollArg as you want,
they will build into a rolling expression. Invalid formats will be ignored.

Example:

dndhelper 1d6 4d4 1d10+1 1D8-1

A rollArg can also be a rollattribute, to build more complex rolling expressions.

rollAttribute string list:

 -roll, hit, dmg : separators, starts a new rolling expressions
 -crit: Critical, doubles all dice ammount
 -spell: Spell, DiceRollResults.String() prints the sum and the sum halved for saves
 -half: Halves the sums, for resistances and such
 -adv: Advantage, rolls each dice twice and drops the lowest
 -dis: Disadvantage, rolls each dice twice and drops the highest
 -drophigh: Drop High, drops the highest result of a DiceRoll
 -droplow: Drop Low, drops the lowest result of a DiceRoll

Examples:

dndhelper hit advantage 1d20+5 dmg 2d6+4 1d4
dndhelper roll 2d10 roll 6d6 roll 1d100
dndhelper spell 6d8
dndhelper droplow 4d6

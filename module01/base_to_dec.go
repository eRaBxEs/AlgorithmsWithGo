package module01

import (
	"math"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {

	// const charset = "0123456789ABCDEF"
	strlen := len(value)
	retv := 0
	baseF := float64(base)
	for k, v := range value {

		runeToInt := 0

		runeToInt, err := strconv.Atoi(runeSelect(string(v)))
		if err != nil {
			return 0
		}
		retv = retv + int(math.Pow(baseF, float64(strlen-(k+1))))*runeToInt
	}

	return retv
}

func runeSelect(st string) string {

	switch st {
	case "A":
		return "10"
	case "B":
		return "11"
	case "C":
		return "12"
	case "D":
		return "13"
	case "E":
		return "14"
	case "F":
		return "15"
	}

	return st
}

func BaseToDec2(value string, base int) int {

	const charset = "0123456789ABCDEF"
	res := 0
	muliplier := 1

	for i := len(value) - 1; i >= 0; i-- { // moving from the rightmost to the leftmost

		index := -1

		for j, char := range charset {

			if char == rune(value[i]) {
				index = j // index count gives the value to rep base value
			}

		}

		res = res + index*muliplier
		muliplier = muliplier * base
	}

	return res
}

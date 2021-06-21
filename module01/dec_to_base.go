package module01

import (
	"fmt"
	"strconv"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {

	var retv string
	var rem int
	for dec > 0 {

		rem = dec % base

		retv = fmt.Sprintf("%s%s", GetBase16Literals(rem), retv)

		dec = dec / base
	}

	return retv
}

func GetBase16Literals(x int) string {

	switch x {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	}
	return strconv.Itoa(x)
}

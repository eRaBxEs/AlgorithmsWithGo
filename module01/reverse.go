package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {

	var res string
	for _, r := range word {
		res = string(r) + res
	}
	return res
	// Makes no provision for other language characters:
	// rev := ""
	// for i := (len(word) - 1); i >= 0; i-- {
	// 	rev = rev + string(word[i])

	// }
	// return rev
}

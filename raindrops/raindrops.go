package raindrops

import "strconv"

func Convert(number int) string {
	var output string
	if number%3 == 0 {
		output += "Pling"
	}
	if number%5 == 0 {
		output += "Plang"
	}
	if number%7 == 0 {
		output += "Plong"
	}
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		numstr := strconv.Itoa(number)
		output += numstr
	}
	return output
}

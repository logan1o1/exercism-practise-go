package resistorcolor

import "strings"

// Colors returns the list of all colors.
func Colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colorArr := Colors()
	colorMap := make(map[string]int)
	for i, col := range colorArr {
		colorMap[col] = i
	}

	for _, c := range colorArr {
		c = strings.ToLower(c)
		if color == c {
			return colorMap[c]
		}
	}
	return 0
}

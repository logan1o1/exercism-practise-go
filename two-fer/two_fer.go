// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// Just checking if the name is empty or not
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

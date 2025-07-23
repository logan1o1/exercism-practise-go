package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](input []T, condition func(T) bool) []T {
	var result []T
	for _, n := range input {
		if condition(n) {
			result = append(result, n)
		}
	}
	return result
}

func Discard[T any](input []T, condition func(T) bool) []T {
	var result []T
	for _, n := range input {
		if !condition(n) {
			result = append(result, n)
		}
	}

	return result
}

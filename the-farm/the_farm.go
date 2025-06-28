package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fdrCalc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	totalFodder, err := fdrCalc.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}
	ftfactor, err := fdrCalc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	fodderPerCow := (totalFodder / float64(numCows)) * ftfactor
	return fodderPerCow, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fdrCalc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	foodPerCow, err := DivideFood(fdrCalc, numCows)
	if err != nil {
		return 0, err
	}
	return foodPerCow, nil
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	numCows int
	message string
}

func (er *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", er.numCows, er.message)
}

func ValidateNumberOfCows(numCows int) error {
	var message string
	switch {
	case numCows < 0:
		message = "there are no negative cows"
	case numCows == 0:
		message = "no cows don't need food"
	default:
		return nil
	}

	return &InvalidCowsError{
		numCows: numCows,
		message: message,
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

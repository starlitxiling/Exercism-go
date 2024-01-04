package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, number int) (float64, error) {
	total, err1 := fodderCalculator.FodderAmount(number)
	if err1 != nil {
		return 0, err1
	}
	factor, err2 := fodderCalculator.FatteningFactor()
	if err2 != nil {
		return 0, err2
	}

	return total * factor / float64(number), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, number int) (float64, error) {
	if number > 0 {
		return DivideFood(fodderCalculator, number)
	}

	return 0, errors.New("invalid number of cows")
}

type MyError struct {
	cows           int
	custom_message string
}

func (m *MyError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", m.cows, m.custom_message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(number int) error {
	if number < 0 {
		return &MyError{
			cows:           number,
			custom_message: "there are no negative cows",
		}
	} else if number == 0 {
		return &MyError{
			cows:           number,
			custom_message: "no cows don't need food",
		}
	} else {
		return nil
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

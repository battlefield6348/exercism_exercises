package thefarm

import (
    "errors"
    "fmt"
)

var ErrInvalidNum = errors.New("invalid number of cows")

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
    amount, err := fodderCalculator.FodderAmount(cows)
    if err != nil {
        return 0, err
    }

    fattening, err := fodderCalculator.FatteningFactor()
    if err != nil {
        return 0, err
    }

    return amount / float64(cows) * fattening, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
    if cows <= 0 {
        return 0, ErrInvalidNum
    }

    return DivideFood(fodderCalculator, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
    if cows < 0 {
        return &MyCustomError{cows, "there are no negative cows"}
    }

    if cows == 0 {
        return &MyCustomError{cows, "no cows don't need food"}
    }
    return nil
}

type MyCustomError struct {
  cows int
  message string
}

func (e *MyCustomError) Error() string {
  return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

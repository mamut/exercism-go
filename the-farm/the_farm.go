package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodderCalculator FodderCalculator, nrOfCows int) (float64, error) {
	fodderAmount, err := fodderCalculator.FodderAmount(nrOfCows)

	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fodderCalculator.FatteningFactor()

	if err != nil {
		return 0, err
	}

	return fodderAmount * fatteningFactor / float64(nrOfCows), nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, nrOfCows int) (float64, error) {
	if nrOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fodderCalculator, nrOfCows)
}

type InvalidCowsError struct {
	nrOfCows int
	message  string
}

func (e InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.nrOfCows, e.message)
}

func ValidateNumberOfCows(nrOfCows int) error {
	if nrOfCows < 0 {
		return &InvalidCowsError{nrOfCows, "there are no negative cows"}
	}

	if nrOfCows == 0 {
		return &InvalidCowsError{nrOfCows, "no cows don't need food"}
	}

	return nil
}

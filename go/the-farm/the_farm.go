package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	numberOfCows int
}

func (sillyNephewError *SillyNephewError) Error() error {
	message := fmt.Sprintf("silly nephew, there cannot be %d cows", sillyNephewError.numberOfCows)
	return errors.New(message)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0, errors.New("division by zero")
	} else if err == ErrScaleMalfunction && fodder >= 0 {
		fodder = fodder * 2
		return fodder / float64(cows), nil
	} else if fodder < 0 && (err == ErrScaleMalfunction || err == nil) {
		return 0, errors.New("negative fodder")
	} else if cows < 0 {
		sillyNephewError := SillyNephewError{
			numberOfCows: cows,
		}
		return 0, sillyNephewError.Error()
	}

	if err != nil {
		return 0, err
	}
	return fodder / float64(cows), nil

}

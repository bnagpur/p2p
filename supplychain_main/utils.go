package main

import (
	"errors"
	"fmt"
)

func LogAndError(errorMessage string) error {
	fmt.Println(errorMessage)
	return errors.New(errorMessage)
}

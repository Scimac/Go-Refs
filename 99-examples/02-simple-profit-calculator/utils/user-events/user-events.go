package userEvents

// Created a custom package here

import (
	"fmt"
	"os"
)

func GetFloat64Input(inputMsg string) float64 {
	fmt.Println(inputMsg)
	var userInput float64
	_, err := fmt.Scanln(&userInput)

	if userInput <= 0 {
		fmt.Println("Input must be a positive number.")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return GetFloat64Input(inputMsg)
	}

	return userInput
}

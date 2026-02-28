package main

import (
	"fmt"
	"io"
	"os"
)

func fileReaderAssignment() {
	cliArgs := os.Args
	fmt.Println("CLI Arguments: ", cliArgs)

	if len(cliArgs) < 2 {
		fmt.Println("Please provide file name to open")
		os.Exit(1)
	}

	fileName := cliArgs[1]
	fmt.Println("File to open: ", fileName)

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}

	n, err := io.Copy(os.Stdout, file)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}
	fmt.Printf("\nTotal bytes read: %d\n", n)
}

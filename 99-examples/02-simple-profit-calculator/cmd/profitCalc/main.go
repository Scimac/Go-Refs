package main

import (
	"fmt"
	"os"
	userEvents "profit-calculator/utils/user-events" // custom public package
)

func main() {
	revenue := userEvents.GetFloat64Input("Enter Revenue: ")
	expenses := userEvents.GetFloat64Input("Enter Expenses: ")
	taxRate := userEvents.GetFloat64Input("Enter Tax Rate (in %): ")

	profit := calculateProfit(revenue, expenses, taxRate)

	fmt.Printf("Net Profit: %.2f\n", profit)
	err := saveToFile("profit.txt", fmt.Sprintf("Net Profit: %.2f\n", profit))
	if err != nil {
		fmt.Println("Error saving to file:", err)
	}
}

func calculateProfit(revenue, expenses, taxRate float64) float64 {

	profitBeforeTax := revenue - expenses
	taxAmount := profitBeforeTax * (taxRate / 100)
	netProfit := profitBeforeTax - taxAmount

	return netProfit
}

func saveToFile(filename string, data string) error {
	f, _ := os.OpenFile(filename, os.O_RDONLY, 0644)

	if f != nil {
		os.Remove(filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

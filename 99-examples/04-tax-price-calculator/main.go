package main

import (
	"fmt"
	dataLoader "tax-price-calculator/data-loader"
	pricespkg "tax-price-calculator/prices"
)

func main() {
	prices := dataLoader.FetchCSVData("./lib/price-data.csv")
	taxRates := dataLoader.FetchCSVData("./lib/tax-data.csv")

	result := make(map[string]*pricespkg.TaxCalculator)

	for _, taxSlab := range taxRates {
		result[fmt.Sprintf("%s-%s", taxSlab["jurisdiction"], taxSlab["type"])] = pricespkg.NewTaxCalculator(prices, taxSlab)
	}

	for key, calculator := range result {
		fmt.Printf("\n\nResults for %s:\n", key)
		calculator.DisplayPrices()
	}

	fmt.Println("\n\nExporting results to CSV file...")
	for key, calculator := range result {
		filePath := fmt.Sprintf("./output/prices-with-tax-%s.csv", key)
		err := calculator.ExportToCSV(filePath)
		if err != nil {
			fmt.Printf("Error exporting to CSV for %s: %v\n", key, err)
		} else {
			fmt.Printf("Exported results to %s successfully.\n", filePath)
		}
	}
}

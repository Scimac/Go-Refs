package prices

import (
	"fmt"
	"sort"
	"strconv"
	dataLoader "tax-price-calculator/data-loader"
)

type TaxCalculator struct {
	TaxRate           map[string]string
	Prices            dataLoader.FileData
	TaxIncludedPrices dataLoader.FileData
	Keys              []string
}

func NewTaxCalculator(prices dataLoader.FileData, taxRate map[string]string) *TaxCalculator {
	tc := &TaxCalculator{
		TaxRate: taxRate,
		Prices:  prices,
	}
	tc.calculatePricesWithTax()
	return tc
}

// CalculatePricesWithTax takes a slice of prices and a tax rate,
// and returns a new slice of prices with the tax applied.
func (tc *TaxCalculator) calculatePricesWithTax() {
	tc.TaxIncludedPrices = dataLoader.FileData{}
	tc.Keys = []string{"item_name", "tax_jurisdiction", "unit_price", "tax_included_price"}

	for _, price := range tc.Prices {
		itemName := price["item_name"]
		taxJurisdiction := tc.TaxRate["jurisdiction"]
		priceValue, _ := strconv.ParseFloat(price["unit_price"], 64)
		taxRate, _ := strconv.ParseFloat(tc.TaxRate["rate"], 64)
		taxIncludedPrice := (priceValue * (1 + taxRate))

		tc.TaxIncludedPrices = append(tc.TaxIncludedPrices, map[string]string{
			"item_name":          itemName,
			"tax_jurisdiction":   taxJurisdiction,
			"unit_price":         fmt.Sprintf("%.2f", priceValue),
			"tax_included_price": fmt.Sprintf("%.2f", taxIncludedPrice),
		})
	}
}

// DisplayPrices prints the original and tax-included prices.
func (tc *TaxCalculator) DisplayPrices() {
	for _, price := range tc.TaxIncludedPrices {
		fmt.Printf("Item: %s | Original Price: $%s | Tax Included Price: $%s \n",
			price["item_name"],
			price["unit_price"],
			price["tax_included_price"],
		)
	}
}

func (tc *TaxCalculator) ExportToCSV(filePath string) error {
	if len(tc.TaxIncludedPrices) == 0 {
		return fmt.Errorf("no data to export")
	}

	// determine keys (keeps deterministic order if tc.Keys is set)
	keys := tc.Keys
	if len(keys) == 0 {
		for k := range tc.TaxIncludedPrices[0] {
			keys = append(keys, k)
		}
		sort.Strings(keys)
	}

	// delegate actual CSV writing to data-loader utility
	return dataLoader.WriteCSV(filePath, keys, tc.TaxIncludedPrices)
}

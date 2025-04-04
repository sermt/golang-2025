package main

import (
	"fmt"

	"example.com/project/filemanager"
	"example.com/project/prices"
)

func main() {
	var taxRates = []float64{0.01, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fileManager := filemanager.NewFileManager("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		priceJob.Process()
	}
}

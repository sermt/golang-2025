package main

import (
	"example.com/project/cmdmanager"
	"example.com/project/prices"
)

func main() {
	var taxRates = []float64{0.01, 0.07, 0.1, 0.15}
	cmd := cmdmanager.NewCMDManager()
	for _, taxRate := range taxRates {

		//fileManager := filemanager.NewFileManager("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		priceJob.Process()
	}
}

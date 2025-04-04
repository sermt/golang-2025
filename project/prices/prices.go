package prices

import (
	"fmt"

	"example.com/project/conversion"
	"example.com/project/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		key := fmt.Sprintf("%.2f", price)
		result[key] = price * (1 + job.TaxRate)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteJSON(job)
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

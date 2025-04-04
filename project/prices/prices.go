package prices

import (
	"fmt"

	"example.com/project/conversion"
	"example.com/project/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
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

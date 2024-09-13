package prices

import (
	"fmt"

	"course.go/price-calculator/conversion"
	"course.go/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	Prices            []float64           `json:"prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(IO iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: IO,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool) {
	err := job.loadPrices()

	if err != nil {
		return
	}

	results := make(map[string]string, len(job.Prices))
	for _, price := range job.Prices {

		// Storing the result into the results map, formating the results in a readable way
		results[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	job.TaxIncludedPrices = results
	job.IOManager.WriteResult(job)
	doneChan <- true
}

// Func to load prices to
func (job *TaxIncludedPriceJob) loadPrices() error {
	readFile, err := job.IOManager.ReadFiles()

	if err != nil {
		return err
	}

	pricesConverted, err := conversion.SliceStringToSliceFloat64(readFile)

	if err != nil {
		return err
	}

	job.Prices = pricesConverted

	return nil
}

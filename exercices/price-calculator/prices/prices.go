package prices

import (
	"fmt"

	"course.go/price-calculator/conversion"
	"course.go/price-calculator/filemanager"
)

const pricesFile = "prices.txt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.loadPrices(pricesFile)
	results := make(map[string]string, len(job.Prices))
	for _, price := range job.Prices {

		// Storing the result into the results map, formating the results in a readable way
		results[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	job.TaxIncludedPrices = results
	filemanager.WriteJSON(job, fmt.Sprintf("result_%.0f.json", job.TaxRate*100))
}

// Func to load prices to
func (job *TaxIncludedPriceJob) loadPrices(pricesFile string) {
	readFile, err := filemanager.ReadFiles(pricesFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	pricesConverted, err := conversion.SliceStringToSliceFloat64(readFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.Prices = pricesConverted
}

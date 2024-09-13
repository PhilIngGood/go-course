package main

import (
	"fmt"

	"course.go/price-calculator/filemanager"
	"course.go/price-calculator/prices"
)

const pricesFile = "prices.txt"

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New(pricesFile, fmt.Sprintf("results_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {

			fmt.Println("An error occured processing a job:", err)
		}
	}

}

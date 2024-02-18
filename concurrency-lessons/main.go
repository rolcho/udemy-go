package main

import (
	// "fmt"
	"fmt"
	"log"

	// "example.com/concurrency/cmdmanager"
	"example.com/concurrency/conversion"
	"example.com/concurrency/filemanager"
	"example.com/concurrency/prices"
)

func main() {

	fm := filemanager.New("taxrates.txt", "")
	taxRatesStr, err := fm.ReadLines()
	if err != nil {
		log.Panic(err.Error())
	}

	taxRates, err := conversion.StringsToFloats(taxRatesStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, taxRate := range taxRates {
		// pm := cmdmanager.New()
		pm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(pm, taxRate)
		if err := priceJob.Process(); err != nil {
			log.Output(0, "ERROR:"+err.Error())
		}
	}

}

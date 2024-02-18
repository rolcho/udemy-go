package main

import (
	"fmt"
	"log"

	"example.com/tax/conversion"
	"example.com/tax/filemanager"
	"example.com/tax/prices"
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
		fmPrice := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fmPrice, taxRate)
		if err := priceJob.Process(fmPrice.InputFilePath); err != nil {
			log.Output(0, "ERROR:"+err.Error())
		}
	}

}

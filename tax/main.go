package main

import (
	"log"

	"example.com/tax/conversion"
	"example.com/tax/filemanager"
	"example.com/tax/prices"
)

func main() {

	taxRatesStr, err := filemanager.ReadLines("taxrates.txt")
	if err != nil {
		log.Panic(err.Error())
	}

	taxRates, err := conversion.StringsToFloats(taxRatesStr)
	if err != nil {
		log.Panic(err.Error())
	}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		if err := priceJob.Process("prices.txt"); err != nil {
			log.Output(0, "ERROR:"+err.Error())
		}
	}

}

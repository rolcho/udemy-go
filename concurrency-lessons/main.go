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

	doneChans := make([]chan bool, len(taxRatesStr))

	taxRates, err := conversion.StringsToFloats(taxRatesStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		// pm := cmdmanager.New()
		pm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(pm, taxRate)
		go priceJob.Process(doneChans[index])

		// if err := priceJob.Process(); err != nil {
		// 	log.Output(0, "ERROR:"+err.Error())
		// }
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}

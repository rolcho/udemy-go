package main

import (
	"fmt"
	"log"

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
	errorChans := make([]chan error, len(taxRatesStr))

	taxRates, err := conversion.StringsToFloats(taxRatesStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		pm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(pm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRatesStr {
		select { //cases wait for one case channel
		case err := <-errorChans[index]:
			if err != nil {
				log.Output(0, "ERROR:"+err.Error())
			}
		case <-doneChans[index]:
			log.Output(0, "Done")
		}
	}
}

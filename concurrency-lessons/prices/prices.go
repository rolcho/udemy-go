package prices

import (
	"errors"
	"fmt"
	"time"

	"example.com/concurrency/conversion"
	"example.com/concurrency/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"taxRate"`
	InputPrices       []float64           `json:"inputPrices"`
	TaxIncludedPrices map[string]string   `json:"taxIncludedPrices"`
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan float64, errorChan chan error) {
	if err := job.LoadPrices(); err != nil {
		errorChan <- err
		return
	}
	if job.TaxRate < 0.1 {
		time.Sleep(500 * time.Millisecond)
		m := fmt.Sprintf(" %.0f percent is too low", job.TaxRate*100)
		errorChan <- errors.New(m)
		return
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
	doneChan <- job.TaxRate
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate:     taxRate,
		IOManager:   iom,
	}
}

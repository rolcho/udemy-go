package prices

import (
	"fmt"

	"example.com/tax/conversion"
	"example.com/tax/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"taxRate"`
	InputPrices       []float64               `json:"inputPrices"`
	TaxIncludedPrices map[string]string       `json:"taxIncludedPrices"`
}

func (job *TaxIncludedPriceJob) LoadPrices(fileName string) error {
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

func (job *TaxIncludedPriceJob) Process(fileName string) error {
	if err := job.LoadPrices(fileName); err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate:     taxRate,
		IOManager:   fm,
	}
}

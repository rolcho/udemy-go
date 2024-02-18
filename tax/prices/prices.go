package prices

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Printf("ERROR %v\n", err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Printf("ERROR %v\n", err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Printf("ERROR %v\n", err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices
	file.Close()

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()

	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		textIncludedPrice := price * (1 + job.TaxRate)
		textIncludedPrice = math.Round(textIncludedPrice*100) / 100
		result[fmt.Sprintf("%.2f", price)] = textIncludedPrice
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}

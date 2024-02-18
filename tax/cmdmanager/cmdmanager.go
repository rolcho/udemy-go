package cmdmanager

import "fmt"

type CmdManager struct {
}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Enter your prices")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CmdManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

package main

import "fmt"

func main() {
	websites := map[string]string{
		"google": "https://google.com",
		"apple":  "https://apple.com",
	}
	websites["microsoft"] = "https://microsoft.com"
	fmt.Println(websites)

	websites["microsoft"] = "https://office.com"
	fmt.Println(websites["microsoft"])

	delete(websites, "microsoft")
	fmt.Println(websites)
}

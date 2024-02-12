package main

import "fmt"

func main() {
	// websites := []string{"https://google.com", "https://apple.com"}
	websites := map[string]string{
		"google": "https://google.com",
		"apple":  "https://apple.com",
	}
	websites["microsoft"] = "https://microsoft.com"

	fmt.Println(websites)

	delete(websites, "microsoft")
	fmt.Println(websites)
}

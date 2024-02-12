package main

import "fmt"

type course map[string]float64

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

	courseRatings := make(course, 3)
	courseRatings["go"] = 4.56
	courseRatings["ionic"] = 4.34

	fmt.Println(courseRatings)

}

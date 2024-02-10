package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age:", age)
	updateToAdultYears(&age)
	fmt.Println("Adult years:", age)
}

func updateToAdultYears(age *int) {
	*age = *age - 18
}

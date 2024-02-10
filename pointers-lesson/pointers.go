package main

import "fmt"

func main() {
	age := 32

	agePointer := &age
	fmt.Println(agePointer)
	fmt.Println(*agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult years: ", adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}

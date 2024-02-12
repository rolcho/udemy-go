package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	var products []Product
	newProduct1 := Product{"id1", "Mac", 1199}
	newProduct2 := Product{"id2", "PC", 1099}
	newProduct3 := Product{"id3", "Mobile", 499}
	products = append(products, newProduct1, newProduct2, newProduct3)

	fmt.Println(products)
	fmt.Println(products[0])

	fmt.Println("iterate array -----------------")
	for i, p := range products {
		fmt.Println(i, p.title)
	}

	fmt.Println("array slice-----------------")
	featuredProducts := products[0:2]
	fmt.Println("basic slice [0] -> [1]:", featuredProducts)
	fmt.Println("fP size:", len(featuredProducts))

	fmt.Println("array slice changes when original updates -----------------")
	products[1].title = "PC laptop"
	featuredProducts = products[1:]
	fmt.Println("fp [1]-> length:", featuredProducts)

	fmt.Println("products | featured | highlighted array sub-slice -----------------")
	highlightedProduct := featuredProducts[:1]
	fmt.Println("Highlighted:", highlightedProduct)

}

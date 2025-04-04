package main

import "fmt"

type Product struct {
	Title string
	Price float64
	Id    string
}

func main() {
	products := []float64{10.5, 10.6, 10.7, 10.8, 10.9}
	fmt.Println(products)

	products = append(products, 10.10)
	fmt.Println(products)

	discountPrices := []float64{5.5, 5.6, 5.7, 5.8, 5.9}

	products = append(products, discountPrices...)
	fmt.Println(products)

}

/* func main() {
	prices := [4]float64{10.5, 20.5, 30.5, 14.5}
	fmt.Println("Prices:", prices[1:4])
	length, capacity := len(prices), cap(prices)

	fmt.Printf("Length: %d, Capacity: %d\n", length, capacity)

	fakePrices := prices[:1]
	length, capacity = len(fakePrices), cap(fakePrices)

	fmt.Println("Fake Prices:", fakePrices)
	fmt.Printf("Length: %d, Capacity: %d\n", length, capacity)

	fakePrices[0] = 100.5
	fmt.Println("Prices:", prices)

} */

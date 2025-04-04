package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(transformNumbers(&numbers, double))
	fmt.Println(transformNumbers(&numbers, square))
}

func transformNumbers(numbers *[]int, transformFunc func(int) int) []int {
	transformedValues := make([]int, len(*numbers))
	for i, number := range *numbers {
		transformedValues[i] = transformFunc(number)
	}
	return transformedValues
}

func double(number int) int {
	return number * 2
}

func square(number int) int {
	return number * number
}

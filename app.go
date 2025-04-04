package main

import (
	"fmt"
	"math"
)

func main() {
	var invesmentAmount float64
	var expectedReturnRate float64
	var years float64

	const inflationRate = 3.5

	fmt.Print("Enter your initial investment amount: ")
	fmt.Scan(&invesmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := float64(invesmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println("Future value after", years, "years: $", futureValue)

	realFutureValue := math.Round(futureValue / math.Pow(1+inflationRate/100, years))
	fmt.Println("Real value after", years, "years with inflation: $", realFutureValue)

}

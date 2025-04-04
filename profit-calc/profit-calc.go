package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue int
	var taxRate float64
	var expenses int

	fmt.Println("Enter revenue:")
	fmt.Scan(&revenue)

	fmt.Println("Enter tax rate (in percentage):")
	fmt.Scan(&taxRate)

	fmt.Println("Enter expenses:")
	fmt.Scan(&expenses)

	ebt := revenue - expenses

	if taxRate == 100 {
		fmt.Println("Error: tax rate cannot be 100%.")
		return
	}

	profit := float64(ebt) * (1 - taxRate/100)

	if profit == 0 {
		fmt.Println("Error: profit is zero, cannot calculate ratio.")
		return
	}

	ratio := float64(ebt) / profit

	// Formatted output for better readability and readability.
	formattedText := fmt.Sprintf(`
    EBT: $%d
    Profit: $%.2f
    Ratio: %.2f`, ebt, profit, ratio)

	fmt.Println(formattedText)

	calculate_future_value()

}

const inflationRate = 2.5

func calculate_future_value() {

	var invesmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Println("Enter invesment amount:")
	fmt.Scan(&invesmentAmount)

	fmt.Println("Expected return rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Enter number of years:")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculate_future_values(invesmentAmount, expectedReturnRate, years)

	fmt.Printf("Future value: $%.2f\n", futureValue)
	fmt.Printf("Future real value: $%.2f\n", futureRealValue)
}

func calculate_future_values(invesmentAmount float64, expectedReturnRate float64, years float64) (futureValue float64, rfv float64) {
	futureValue = invesmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = futureValue / math.Pow(1+inflationRate/100, years)
	return
}

package main

import (
	"fmt"

	"example.com/bank/file_handler"
)

func main() {

	for {
		presentOptions()
		var choice int
		fileName := "balance.txt"
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		fmt.Println("You chose option", choice)
		accountBalance, error := file_handler.GetValueFromFile(fileName)

		if error != nil && choice != 4 {
			fmt.Println("Error reading from file:", error)
			//return
			panic("Error reading from file")
		}
		switch choice {
		case 1:
			fmt.Println("Checking balance...")
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				continue
			}
			accountBalance += depositAmount
			file_handler.WriteValueToFile(accountBalance, fileName)
			fmt.Printf("Deposit successful! New balance is: $%.2f\n", accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("Enter the withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			} else {
				accountBalance -= withdrawalAmount
				file_handler.WriteValueToFile(accountBalance, fileName)
				fmt.Printf("Withdrawal successful! New balance is: $%.2f\n", accountBalance)
			}
		case 4:
			fmt.Println("Thank you for using Go Bank! Goodbye! :)")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Println("Welcome to go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		fmt.Println("You chose option", choice)
		accountBalance, error := getBalanceFromFile()
		/* 	wantsToCheckBalance := choice == 1
		wantsToDeposit := choice == 2
		wantsToWithdraw := choice == 3
		wantsToExit := choice == 4

		if wantsToCheckBalance {
			fmt.Println("Checking balance...")
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)

		} else if wantsToDeposit {
			var depositAmount float64
			fmt.Print("Enter the deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Deposit successful! New balance is: $%.2f\n", accountBalance)
		} else if wantsToWithdraw {

			var withdrawalAmount float64
			fmt.Print("Enter the withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			} else {
				accountBalance -= withdrawalAmount
				fmt.Printf("Withdrawal successful! New balance is: $%.2f\n", accountBalance)
			}
		} else if wantsToExit {
			fmt.Println("Exiting go Bank...")

			break
		} else if choice < 1 || choice > 4 {
			fmt.Println("Invalid choice. Please try again.")
		} */

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
			writeBalanceToFile(accountBalance)
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
				writeBalanceToFile(accountBalance)
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

func writeBalanceToFile(balance float64) error {
	file, err := os.OpenFile("balance.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%.2f", balance))
	return err
}

func getBalanceFromFile() (float64, error) {
	file, err := os.Open("balance.txt")
	if err != nil {
		return 0.0, err
	}
	defer file.Close()

	var balance float64
	_, err = fmt.Fscanf(file, "%f", &balance)
	return balance, err
}

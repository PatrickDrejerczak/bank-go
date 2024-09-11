package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to read balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse balance")
	}

	return balance, nil
}

func writeBlanaceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var choice int
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println(("----------------"))
	}

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1
		// wantsDepositMoney := choice == 2
		// wantsWithdrawMoney := choice == 3
		// wantsToExit := choice == 4

		// if choice < 1 || choice > 4 {
		// 	fmt.Println("Invalid choice")
		// }

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance, "$")
		case 2:
			fmt.Print("Enter amount to deposit: ")

			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Value must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance is", accountBalance, "$")
			writeBlanaceToFile(accountBalance)
		case 3:
			fmt.Print("Enter amount to withdraw: ")

			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Value must be greater than 0")
				continue
			}

			if accountBalance-withdrawAmount < 0 {
				fmt.Println("You don't have enough balance to withdraw this amount")
			} else {
				accountBalance -= withdrawAmount
			}

			fmt.Println("Your new balance is", accountBalance, "$")
			writeBlanaceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank")
			return
		}

		// 	if wantsCheckBalance {
		// 		fmt.Println("Your balance is ", accountBalance, "$")
		// 	} else if wantsDepositMoney {
		// 		fmt.Print("Enter amount to deposit: ")

		// 		var depositAmount float64
		// 		fmt.Scan(&depositAmount)

		// 		if depositAmount <= 0 {
		// 			fmt.Println("Invalid amount. Value must be greater than 0")
		// 			continue
		// 		}

		// 		accountBalance += depositAmount
		// 		fmt.Println("Your new balance is", accountBalance, "$")
		// 	} else if wantsWithdrawMoney {
		// 		fmt.Print("Enter amount to withdraw: ")

		// 		var withdrawAmount float64
		// 		fmt.Scan(&withdrawAmount)

		// 		if withdrawAmount <= 0 {
		// 			fmt.Println("Invalid amount. Value must be greater than 0")
		// 			continue
		// 		}

		// 		if accountBalance-withdrawAmount < 0 {
		// 			fmt.Println("You don't have enough balance to withdraw this amount")
		// 			continue
		// 		} else {
		// 			accountBalance -= withdrawAmount
		// 		}

		// 		fmt.Println("Your new balance is", accountBalance, "$")
		// 	} else if wantsToExit {
		// 		fmt.Println("Goodbye!")
		// 		break
		// 	}
	}
}

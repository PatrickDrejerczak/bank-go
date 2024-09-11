package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var choice int
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println(("----------------"))
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())

	for {
		presentOptions()
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank")
			return
		}
	}
}

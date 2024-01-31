package main

import (
	"bankingapp/accountopen"
	"fmt"
)

func main() {
	fmt.Println("Welcome to banking")
	acc1 := accountopen.OpenAccount("Lakshmi", "password@123", 1000.0)
	acc2 := accountopen.OpenAccount("Priya", "priya@123", 550.0)
	fmt.Println("Account successfully opened!", acc1)
	fmt.Println("Account successfully opened!", acc2)
	fmt.Println("Initial balance from account1:", acc1.Balance)
	acc1.AddAmount(500.0)
	fmt.Println("Balance after adding money:", acc1.Balance)
	acc1.DeductAmount(50.0)
	fmt.Println("Balance after DeductAmount:", acc1.Balance)

	amount := 100.0
	err := accountopen.Transfer(&acc1, &acc2, amount)
	if err != nil {
		fmt.Println("Transfer failed:", err)
	} else {
		fmt.Println("Transfer successful!")
		fmt.Println("New balances:")
		fmt.Println("Final amount in Account 1:", acc1.Balance)
		fmt.Println("Final amount in Account 2:", acc2.Balance)
	}

	totalAccounts := accountopen.GetTotalAccounts()
	fmt.Println("Total accounts:", totalAccounts)

	messege := accountopen.CloseAccount("Lakshmi")
	if err != nil {
		fmt.Println("Error closing account:", messege)
	} else {
		fmt.Println("Account  closed successfully!")
	}

	// Get and print total number of accounts after closing
	fmt.Println("Total accounts after closing:", accountopen.GetTotalAccounts())
}

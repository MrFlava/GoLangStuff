package main

import "fmt"

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main(){

	myAccounts := []IBankAccount {
		NewWellsFargo(),
		NewBitcoinAccount(),
	}


	for _, account := range myAccounts {
		account.Deposit(500)
		account.Withdraw(70)
		//  err != nil {
		// 	fmt.Printf("ERR: %d\n", err)
		// }

		balance := account.GetBalance()
		fmt.Printf("balance = %d\n", balance)
	}
	// wf := NewWellsFargo()

	// wf.Deposit(200)
	// wf.Deposit(100)
	// wf.Deposit(3201)

	// currentBalance := wf.GetBalance()

	// fmt.Printf("wF balance: %d\n", currentBalance)
}
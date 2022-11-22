package main

import (
	"banco/accounts"
	"banco/client"
	"fmt"
)

func payBill(account verifyAccount, amount float64) {
	account.Withdraw(amount)
}

type verifyAccount interface {
	Withdraw(withdrawnAmount float64) string
}

func main() {

	accountDenis := accounts.SavingsAccount{
		Bearer:        client.Bearer{Name: "Denis", CPF: "123.456.789-10", Profession: "Developer"},
		Agency:        123,
		AccountNumber: 123456,
		Operation:     0,
	}

	accountDenis.Deposit(100)

	payBill(&accountDenis, 60)

	fmt.Println("Balance:", accountDenis.ViewBalance())

}

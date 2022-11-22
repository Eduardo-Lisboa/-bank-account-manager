package accounts

import (
	"banco/client"
)

type SavingsAccount struct {
	Bearer        client.Bearer
	Agency        int
	AccountNumber int
	Operation     int
	balance       float64
}

func (c *SavingsAccount) Withdraw(withdrawnAmount float64) string {
	canWithdraw := withdrawnAmount > 0 && withdrawnAmount <= c.balance
	if canWithdraw {
		c.balance -= withdrawnAmount
		return "Withdrawn successfully"
	} else {
		return "Insufficient balance"
	}

}
func (c *SavingsAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "Deposit successfully", c.balance
	} else {
		return "The value must be greater than zero", c.balance
	}

}

func (c *SavingsAccount) ViewBalance() float64 {
	return c.balance
}

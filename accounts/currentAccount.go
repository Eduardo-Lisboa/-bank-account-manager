package accounts

import (
	"banco/client"
)

type CurrentAccount struct {
	Bearer        client.Bearer
	Agency        int
	AccountNumber int
	balance       float64
}

func (c *CurrentAccount) Withdraw(withdrawnAmount float64) string {
	canWithdraw := withdrawnAmount > 0 && withdrawnAmount <= c.balance
	if canWithdraw {
		c.balance -= withdrawnAmount
		return "Withdrawn successfully"
	} else {
		return "Insufficient balance"
	}

}
func (c *CurrentAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "Deposit successfully", c.balance
	} else {
		return "The value must be greater than zero", c.balance
	}

}
func (c *CurrentAccount) transfer(transferAmount float64, accountDestination *CurrentAccount) bool {
	if transferAmount < c.balance && transferAmount > 0 {
		c.balance -= transferAmount
		accountDestination.Deposit(transferAmount)
		return true
	} else {
		return false
	}

}

func (c *CurrentAccount) ViewBalance() float64 {
	return c.balance
}

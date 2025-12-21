package main

import "fmt"

type Account struct {
	owner   string
	number  int
	branch  int
	balance float64
}

func (c *Account) Withdraw(withdrawAmount float64) string {
	canWithdraw := withdrawAmount > 0 && withdrawAmount <= c.balance

	if canWithdraw {
		c.balance -= withdrawAmount
		return "Saque realizado com sucesso"
	}
	return "Saldo insuficiente"
}

func main() {
	// Different ways to create a new account:
	acc1 := Account{
		owner:   "Fulano",
		number:  123456,
		branch:  1234,
		balance: 1000.00,
	}

	var acc2 *Account
	acc2 = new(Account)
	acc2.owner = "Ciclano"
	acc2.number = 654321
	acc2.branch = 4321
	acc2.balance = 2000.00

	acc3 := Account{"Sicrano", 987654, 3219, 3000.00}

	fmt.Println(acc1)
	fmt.Println(*acc2)
	fmt.Println(acc3)

	result := acc1.Withdraw(200.0)
	fmt.Println(result)
	fmt.Println(acc1.balance)
}

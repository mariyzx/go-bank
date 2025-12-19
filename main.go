package main

import "fmt"

type Account struct {
	owner   string
	number  int
	branch  int
	balance float64
}

func main() {
	conta1 := Account{
		owner:   "Fulano",
		number:  123456,
		branch:  1234,
		balance: 1000.00,
	}

	conta2 := Account{
		owner:   "Ciclano",
		number:  654321,
		branch:  4321,
		balance: 2000.00,
	}

	fmt.Println(conta1)
	fmt.Println(conta2)
}

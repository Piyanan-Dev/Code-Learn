package main

import "fmt"

type Account struct {
	Name   string
	Blance float64
}

func (a *Account) Deposit(n float64) {
	a.Blance += n
}

func (a *Account) Withdraw(n float64) {
	if a.Blance >= n {
		a.Blance -= n
	} else {
		fmt.Print("insufficient")
	}
}

func main() {
	acc := []Account{
		{Name: "A", Blance: 0},
		{Name: "B", Blance: 500},
		{Name: "C", Blance: 50},
	}

	for _, v := range acc {

		v.Deposit(0)

	}

}

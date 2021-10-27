package main

import (
	"fmt"
	"log"

	"github.com/kyunghyun/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(20)
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}

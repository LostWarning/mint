package transact

import (
	"fmt"
)

type Account interface {
	Credit(uint64) uint64
	Debit(uint64) uint64
}

func Deposit(acc Account) {
	var amount uint64
	fmt.Print("Deposit amount : ")
	fmt.Scanf("%d", &amount)
	acc.Credit(amount)
}

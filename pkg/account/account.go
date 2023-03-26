package account

type savingsAccount struct {
	number  string
	balance uint64
	intrest int8
}

func (acc *savingsAccount) Credit(amount uint64) uint64 {
	acc.balance += amount
	return acc.balance
}

func (acc *savingsAccount) Debit(amount uint64) uint64 {
	acc.balance -= amount
	return acc.balance
}

func (acc *savingsAccount) Balance() uint64 {
	return acc.balance
}

func NewSavingsAccount(number string) savingsAccount {
	var acc savingsAccount
	acc.number = number
	acc.balance = 0
	acc.intrest = 0
	return acc
}

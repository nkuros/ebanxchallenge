package model

type Account struct {
	Id string
	Balance int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) {
	a.Balance -= amount
}

func (a *Account) Transfer(amount int, target *Account) {
	a.Withdraw(amount)
	target.Deposit(amount)
}

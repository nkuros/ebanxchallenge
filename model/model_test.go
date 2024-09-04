package model

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	a := Account{Id: "1", Balance: 0}
	a.Deposit(100)
	if a.Balance != 100 {
		t.Errorf("Expected 100, got %d", a.Balance)
	}
}

func TestWithdraw(t *testing.T) {
	a := Account{Id: "1", Balance: 100}
	a.Withdraw(50)
	if a.Balance != 50 {
		t.Errorf("Expected 50, got %d", a.Balance)
	}
}

func TestTransfer(t *testing.T) {
	a := Account{Id: "1", Balance: 100}
	target := Account{Id: "2", Balance: 0}
	a.Transfer(50, &target)
	if a.Balance != 50 {
		t.Errorf("Expected 50, got %d", a.Balance)
	}
	if target.Balance != 50 {
		t.Errorf("Expected 50, got %d", target.Balance)
	}
}


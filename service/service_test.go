package service

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/nkuros/ebanxchallenge/entity"
	"github.com/nkuros/ebanxchallenge/service/mock"
	
)

func TestGetAccount(t *testing.T) {
	ctrl :=	gomock.NewController(t)
	service := mock_service.NewMockAccountService(ctrl)
	service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true)

	account, exists := service.GetAccount("1")
	if exists == false {
		t.Errorf("Expected true, got %t", exists)
	}
	if account.Balance != 100 {
		t.Errorf("Expected 100, got %d", account.Balance)
	}

}

func TestAddAccount(t *testing.T) {
	ctrl :=	gomock.NewController(t)
	service := mock_service.NewMockAccountService(ctrl)
	service.EXPECT().AddAccount("1", 100)
	service.AddAccount("1", 100)
	service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true)
	expected, exists := service.GetAccount("1")
	if exists == true {
		if expected.Balance != 100 {
			t.Errorf("Expected 110, got %d", expected.Balance)
		}
	} else {
		t.Errorf("Account not created")
	}
}

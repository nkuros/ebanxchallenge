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
func TestGetAccountNotFound(t *testing.T) {
	ctrl :=	gomock.NewController(t)
	service := mock_service.NewMockAccountService(ctrl)
	service.EXPECT().GetAccount("1").Return(nil, false)

	account, exists := service.GetAccount("1")
	if exists == true {
		t.Errorf("Expected false, got %t", exists)
	}
	if account != nil {
		t.Errorf("Expected nil, got %v", account)
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

func TestAddAccountAlreadyExists(t *testing.T) {
	ctrl :=	gomock.NewController(t)
	service := mock_service.NewMockAccountService(ctrl)
	service.EXPECT().AddAccount("1", 100).Return(&entity.Account{Id: "1", Balance: 100}, true)
	service.AddAccount("1", 100)

	service.EXPECT().AddAccount("1", 100).Return(&entity.Account{Id: "1", Balance: 100}, false)
	service.AddAccount("1", 100)
}

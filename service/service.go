package service

import (
	"github.com/nkuros/ebanxchallenge/database"
	"github.com/nkuros/ebanxchallenge/entity"
)

//go:generate mockgen -source=service.go -destination=mock/service.go

type accountService struct {
}

type AccountService interface {
	GetAccount(originId string) (account *entity.Account, exists bool)
	AddAccount(originId string, Balance int) (account *entity.Account, created bool)
}

func NewAccountService() AccountService {
	return &accountService{}
}

func (s *accountService) GetAccount(originId string) (*entity.Account, bool) {
	account, exists := database.Accounts[originId]
	return account, exists
}

func (s *accountService) AddAccount(originId string, Balance int) (*entity.Account, bool) {
	account, exists := database.Accounts[originId]
	if exists {
		return account, false
	}
	newAccount := &entity.Account{Id: originId, Balance: Balance}
	database.Accounts[originId] = newAccount

	return newAccount, true
}

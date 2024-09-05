package service

import (
	"github.com/nkuros/ebanxchallenge/database"
	"github.com/nkuros/ebanxchallenge/entity"
)

//go:generate mockgen -source=service.go -destination=mock/service.go

type accountService struct {
}

type AccountService interface {
	GetAccount(originId string) (*entity.Account, bool)
	AddAccount(originId string, Balance int)
}

func NewAccountService() AccountService {
	return &accountService{}
}

func (s *accountService) GetAccount(originId string) (*entity.Account, bool) {
	account, exists := database.Accounts[originId]
	if exists == false {
		return nil, false
	}

	return account, true
}

func (s *accountService) AddAccount(originId string, Balance int) {
	database.Accounts[originId] = &entity.Account{Id: originId, Balance: Balance}
}

package controller

import (
	"fmt"
	"strconv"

	"github.com/nkuros/ebanxchallenge/database"
	"github.com/nkuros/ebanxchallenge/entity"
	"github.com/nkuros/ebanxchallenge/errors"
	"github.com/nkuros/ebanxchallenge/service"
)

//go:generate mockgen -source=controller.go -destination=mock/controller.go

type accountController struct {
	accountService service.AccountService
}

type AccountController interface {
	GetBalanceController(originId string) (string, error)
	PostDepositEventController(originId string, amount int) (string, error)
	PostWithdrawEventController(originId string, amount int) (string, error)
	PostTransferEventController(originId string, targetId string, amount int) (string, error)
	DeleteAllAccountsController()
}

func NewAccountController(accountService service.AccountService) AccountController {
	return &accountController{accountService}
}

func (s *accountController) GetBalanceController(originId string) (string, error) {
	if originId == "" {
		return "0", errors.ErrMissingOriginId
	}
	account, exists := s.accountService.GetAccount(originId)
	if exists == false {
		return "0", errors.ErrOriginAccountNotFound
	}

	return strconv.Itoa(account.Balance), nil
}

func (s *accountController) PostDepositEventController(originId string, amount int) (string, error) {
	if originId == "" {
		return "0", errors.ErrMissingOriginId
	}

	if amount < 0 {
		err := errors.ErrInvalidAmountFormat
		return "0", err
	}
	account, exists := s.accountService.GetAccount(originId)
	if exists == false {
		s.accountService.AddAccount(originId, amount)
		res := fmt.Sprintf("{\"destination\": {\"id\":\"%s\", \"balance\":%d}}", originId, amount)
		return res, nil
	}

	account.Deposit(amount)
	res := fmt.Sprintf("{\"destination\": {\"id\":\"%s\", \"balance\":%d}}", account.Id, account.Balance)
	return res, nil
}

func (s *accountController) PostWithdrawEventController(originId string, amount int) (string, error) {
	if originId == "" {
		return "0", errors.ErrMissingOriginId
	}
	if amount < 0 {
		err := errors.ErrInvalidAmountFormat
		return "0", err
	}

	account, exists := s.accountService.GetAccount(originId)

	if exists == false {
		return "0", errors.ErrOriginAccountNotFound
	}
	if account.Balance < amount {
		return "0", errors.ErrInsufficientFunds
	}

	account.Withdraw(amount)
	res := fmt.Sprintf("{\"origin\": {\"id\":\"%s\", \"balance\":%d}}", account.Id, account.Balance)
	return res, nil
}

func (s *accountController) PostTransferEventController(originId string, targetId string, amount int) (string, error) {
	if originId == "" {
		return "0", errors.ErrMissingOriginId
	}
	if targetId == "" {
		return "0", errors.ErrMissingDestinationId
	}
	if originId == targetId {
		return "0", errors.ErrInvalidDestinationId
	}

	if amount < 0 {
		err := errors.ErrInvalidAmountFormat
		return "0", err
	}
	originAccount, exists := s.accountService.GetAccount(originId)

	if exists == false {
		return "0", errors.ErrOriginAccountNotFound
	}
	targetAccount, exists := s.accountService.GetAccount(targetId)
	if exists == false {
		return "0", errors.ErrTargetAccountNotFound
	}
	if originAccount.Balance < amount {
		return "0", errors.ErrInsufficientFunds
	}

	originAccount.Transfer(amount, targetAccount)
	res := fmt.Sprintf("{\"origin\": {\"id\":\"%s\", \"balance\":%d}, \"destination\": {\"id\":\"%s\", \"balance\":%d}}", originAccount.Id, originAccount.Balance, targetAccount.Id, targetAccount.Balance)
	return res, nil
}

func (s *accountController) DeleteAllAccountsController() {
	database.Accounts = make(map[string]*entity.Account)
}

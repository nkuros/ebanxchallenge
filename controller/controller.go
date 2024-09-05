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

func (c *accountController) GetBalanceController(originId string) (string, error) {
	if originId == "" {
		return "0", errors.ErrMissingOriginId
	}
	account, exists := c.accountService.GetAccount(originId)
	if exists == false {
		return "0", errors.ErrOriginAccountNotFound
	}

	return strconv.Itoa(account.Balance), nil
}

func (c *accountController) PostDepositEventController(destinationId string, amount int) (string, error) {
	if destinationId == "" {
		return "0", errors.ErrMissingOriginId
	}

	if amount < 0 {
		err := errors.ErrInvalidAmountFormat
		return "0", err
	}
	account, exists := c.accountService.GetAccount(destinationId)
	if exists == false {
		_, created := c.accountService.AddAccount(destinationId, amount)
		if created == false {
			return "0", errors.ErrAccountCreationFailed
		}
		res := fmt.Sprintf("{\"destination\": {\"id\":\"%s\", \"balance\":%d}}", destinationId, amount)
		return res, nil
	}

	account.Deposit(amount)
	res := fmt.Sprintf("{\"destination\": {\"id\":\"%s\", \"balance\":%d}}", account.Id, account.Balance)
	return res, nil
}

func (c *accountController) PostWithdrawEventController(originId string, amount int) (string, error) {
	if originId == "" {
		return "0", errors.ErrMissingOriginId
	}
	if amount < 0 {
		err := errors.ErrInvalidAmountFormat
		return "0", err
	}

	account, exists := c.accountService.GetAccount(originId)

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

func (c *accountController) PostTransferEventController(originId string, targetId string, amount int) (string, error) {
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
	originAccount, exists := c.accountService.GetAccount(originId)
	if exists == false {
		return "0", errors.ErrOriginAccountNotFound
	}
	targetAccount, exists := c.accountService.GetAccount(targetId)
	if exists == false {
		targetAccount, exists = c.accountService.AddAccount(targetId, 0)
		if exists == false {
			return "0", errors.ErrAccountCreationFailed
		}
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

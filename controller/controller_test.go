package controller

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/nkuros/ebanxchallenge/database"
	"github.com/nkuros/ebanxchallenge/entity"
	"github.com/nkuros/ebanxchallenge/errors"

	mock_service "github.com/nkuros/ebanxchallenge/service/mock"
)

func TestGetBalanceController(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)
	service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true).Times(1)

	balance, _ := controller.GetBalanceController("1")

	if balance != "100" {
		t.Errorf("Expected 100, got %s", balance)
	}
}
func TestGetBalanceControllerMissingId(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	_, err := controller.GetBalanceController("")

	if err != errors.ErrMissingOriginId {
		t.Errorf("Expected ErrMissingOriginId, got %s", err)
	}
}
func TestGetBalanceControllerNotFound(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)
	service.EXPECT().GetAccount("2").Return(nil, false).Times(1)

	_, err := controller.GetBalanceController("2")

	if err != errors.ErrOriginAccountNotFound {
		t.Errorf("Expected ErrOriginAccountNotFound, got %s", err)
	}
}

func TestDeposit(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)
	service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true).Times(1)

	res, _ := controller.PostDepositEventController("1", 50)

	if res != "{\"destination\": {\"id\":\"1\", \"balance\":150}}" {
		t.Errorf("Expected {\"destination\": {\"id\":\"1\", \"balance\":150}}, got %s", res)
	}
}

func TestDepositMissingId(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	_, err := controller.PostDepositEventController("", 100)

	if err != errors.ErrMissingOriginId {
		t.Errorf("Expected ErrMissingOriginId, got %s", err)
	}
}

func TestDepositInvalidAmount(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	_, err := controller.PostDepositEventController("1", -100)

	if err != errors.ErrInvalidAmountFormat {
		t.Errorf("Expected ErrInvalidAmountFormat, got %s", err)
	}
}

func TestDepositZero(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)
	service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 10}, true)

	res, _ := controller.PostDepositEventController("1", 0)

	if res != "{\"destination\": {\"id\":\"1\", \"balance\":10}}" {
		t.Errorf("Expected {\"destination\": {\"id\":\"1\", \"balance\":10}}, got %s", res)
	}
}

func TestWithdraw(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)
	service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true)

	res, _ := controller.PostWithdrawEventController("1", 49)

	if res != "{\"origin\": {\"id\":\"1\", \"balance\":51}}" {
		t.Errorf("Expected {\"origin\": {\"id\":\"1\", \"balance\":51}}, got %s", res)
	}
}

func TestWithdrawMissingId(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	_, err := controller.PostWithdrawEventController("", 100)

	if err != errors.ErrMissingOriginId {
		t.Errorf("Expected ErrMissingOriginId, got %s", err)
	}
}

func TestWithdrawInvalidAmount(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	_, err := controller.PostWithdrawEventController("1", -100)

	if err != errors.ErrInvalidAmountFormat {
		t.Errorf("Expected ErrInvalidAmountFormat, got %s", err)
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true)
	_, err := controller.PostWithdrawEventController("1", 101)

	if err != errors.ErrInsufficientFunds {
		t.Errorf("Expected ErrInsufficientFunds, got %s", err)
	}
}

func TestWithdrawZero(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)
	service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true).Times(1)

	res, _ := controller.PostWithdrawEventController("1", 0)

	if res != "{\"origin\": {\"id\":\"1\", \"balance\":100}}" {
		t.Errorf("Expected {\"origin\": {\"id\":\"1\", \"balance\":100}}, got %s", res)
	}
}
func TestTransfer(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	gomock.InOrder(
		service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true).Times(1),
		service.EXPECT().GetAccount("2").Return(&entity.Account{Id: "2", Balance: 0}, true).Times(1),
	)


	res, _ := controller.PostTransferEventController("1", "2", 50)

	if res != "{\"origin\": {\"id\":\"1\", \"balance\":50}, \"destination\": {\"id\":\"2\", \"balance\":50}}" {
		t.Errorf("Expected {\"origin\": {\"id\":\"1\", \"balance\":50}, \"destination\": {\"id\":\"2\", \"balance\":50}}, got %s", res)
	}
}

func TestTransferMissingOriginId(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	_, err := controller.PostTransferEventController("", "2", 50)

	if err != errors.ErrMissingOriginId {
		t.Errorf("Expected ErrMissingOriginId, got %s", err)
	}
}

func TestTransferMissingTargetId(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	_, err := controller.PostTransferEventController("1", "", 50)

	if err != errors.ErrMissingDestinationId {
		t.Errorf("Expected ErrMissingTargetId, got %s", err)
	}
}

func TestTransferInvalidAmount(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	_, err := controller.PostTransferEventController("1", "2", -50)

	if err != errors.ErrInvalidAmountFormat {
		t.Errorf("Expected ErrInvalidAmountFormat, got %s", err)
	}
}

func TestTransferInsufficientFunds(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)

	gomock.InOrder(
		service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true).Times(1),
		service.EXPECT().GetAccount("2").Return(&entity.Account{Id: "2", Balance: 0}, true).Times(1),
	)


	_, err := controller.PostTransferEventController("1", "2", 101)

	if err != errors.ErrInsufficientFunds {
		t.Errorf("Expected ErrInsufficientFunds, got %s", err)
	}
}

func TestTransferZero(t *testing.T) {
	service := mock_service.NewMockAccountService(gomock.NewController(t))
	controller := NewAccountController(service)
	
	gomock.InOrder(
		service.EXPECT().GetAccount("1").Return(&entity.Account{Id: "1", Balance: 100}, true).Times(1),
		service.EXPECT().GetAccount("2").Return(&entity.Account{Id: "2", Balance: 0}, true).Times(1),
	)


	res, _ := controller.PostTransferEventController("1", "2", 0)

	if res != "{\"origin\": {\"id\":\"1\", \"balance\":100}, \"destination\": {\"id\":\"2\", \"balance\":0}}" {
		t.Errorf("Expected {\"origin\": {\"id\":\"1\", \"balance\":100}, \"destination\": {\"id\":\"2\", \"balance\":0}}, got %s", res)
	}
}

func TestDeleteAllAccountsController(t *testing.T) {
	database.Accounts = make(map[string]*entity.Account)            //in a real scenario, this would be generated by a mock factory
	database.Accounts["1"] = &entity.Account{Id: "1", Balance: 100} //in a real scenario, this would be mocked
	service := mock_service.NewMockAccountService(gomock.NewController(t))

	controller := NewAccountController(service)

	controller.DeleteAllAccountsController()
	if len(database.Accounts) != 0 {
		t.Errorf("Expected 0, got %d", len(database.Accounts))
	}
}
func TestDeleteAllAccountsControllerEmpty(t *testing.T) {
	database.Accounts = make(map[string]*entity.Account) //in a real scenario, this would be generated by a mock factory
	service := mock_service.NewMockAccountService(gomock.NewController(t))

	controller := NewAccountController(service)

	controller.DeleteAllAccountsController()
	if len(database.Accounts) != 0 {
		t.Errorf("Expected 0, got %d", len(database.Accounts))
	}
}

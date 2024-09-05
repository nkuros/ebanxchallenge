package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	controller_mock "github.com/nkuros/ebanxchallenge/controller/mock"
	"github.com/nkuros/ebanxchallenge/errors"

	"go.uber.org/mock/gomock"
)

func TestGetRoot(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	mock_controller := controller_mock.NewMockAccountController(gomock.NewController(t))
	mock_handler := NewAccountHandler(mock_controller)

	handler := http.HandlerFunc(mock_handler.GetRootHandler)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "EBANX Challenge Root\n"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestGetBalance(t *testing.T) {
	req, err := http.NewRequest("GET", "/balance?account_id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	mock_controller := controller_mock.NewMockAccountController(gomock.NewController(t))
	mock_controller.EXPECT().GetBalanceController("1").Return("100", nil)
	mock_handler := NewAccountHandler(mock_controller)

	handler := http.HandlerFunc(mock_handler.GetBalanceHandler)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := fmt.Sprintf("%d %s", http.StatusOK, "100")
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestGetBalanceInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/balance?account_id=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	mock_controller := controller_mock.NewMockAccountController(gomock.NewController(t))
	mock_controller.EXPECT().GetBalanceController("invalid").Return("", errors.ErrInvalidOriginId)
	mock_handler := NewAccountHandler(mock_controller)

	handler := http.HandlerFunc(mock_handler.GetBalanceHandler)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	expected := "404 0"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestGetBalanceMissing(t *testing.T) {
	req, err := http.NewRequest("GET", "/balance", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	mock_controller := controller_mock.NewMockAccountController(gomock.NewController(t))
	mock_handler := NewAccountHandler(mock_controller)
	mock_controller.EXPECT().GetBalanceController("").Return("", errors.ErrMissingOriginId)

	handler := http.HandlerFunc(mock_handler.GetBalanceHandler)

	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	expected := "404 0"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestPostEvent(t *testing.T) {
	req, err := http.NewRequest("POST", "/event", strings.NewReader("{\"type\":\"deposit\", \"origin\":\"1\", \"amount\":100}"))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	mock_controller := controller_mock.NewMockAccountController(gomock.NewController(t))
	mock_handler := NewAccountHandler(mock_controller)

	mock_controller.EXPECT().PostDepositEventController("1", 100).Return(fmt.Sprintf("{\"destination\": {\"id\":\"1\", \"balance\":100}}"), nil)

	handler := http.HandlerFunc(mock_handler.PostEventHandler)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := fmt.Sprintf("%d {\"destination\": {\"id\":\"1\", \"balance\":100}}", http.StatusCreated)
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestPostEventInvalid(t *testing.T) {
	req, err := http.NewRequest("POST", "/event", strings.NewReader("{\"type\":\"invalid\", \"origin\":\"1\", \"amount\":100}"))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	mock_controller := controller_mock.NewMockAccountController(gomock.NewController(t))
	mock_handler := NewAccountHandler(mock_controller)

	handler := http.HandlerFunc(mock_handler.PostEventHandler)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	expected := "404 0"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestPostEventMissingBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/event", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	mock_controller := controller_mock.NewMockAccountController(gomock.NewController(t))
	mock_handler := NewAccountHandler(mock_controller)
	
	handler := http.HandlerFunc(mock_handler.PostEventHandler)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	expected := "404 0"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestPostEventMissingField(t *testing.T) {
	req, err := http.NewRequest("POST", "/event", strings.NewReader("\"origin\":\"1\", \"amount\":100}"))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	mock_controller := controller_mock.NewMockAccountController(gomock.NewController(t))
	mock_handler := NewAccountHandler(mock_controller)
	
	handler := http.HandlerFunc(mock_handler.PostEventHandler)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	expected := "404 0"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}


func TestPostDelete(t *testing.T) {
	req, err := http.NewRequest("POST", "/delete", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	mock_controller := controller_mock.NewMockAccountController(gomock.NewController(t))
	mock_handler := NewAccountHandler(mock_controller)
	mock_controller.EXPECT().DeleteAllAccountsController().Return()

	handler := http.HandlerFunc(mock_handler.PostDeleteHandler)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "200"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

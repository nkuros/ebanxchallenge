package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/nkuros/ebanxchallenge/constants"
	"github.com/nkuros/ebanxchallenge/controller"
	"github.com/nkuros/ebanxchallenge/errors"
	"github.com/nkuros/ebanxchallenge/model"
)

type accountHandler struct {
	accountController controller.AccountController
}

type AccountHandler interface {
	GetRootHandler(w http.ResponseWriter, req *http.Request)
	GetBalanceHandler(w http.ResponseWriter, r *http.Request)
	PostEventHandler(w http.ResponseWriter, req *http.Request)
	PostDeleteHandler(w http.ResponseWriter, req *http.Request)
}

func NewAccountHandler(accountController controller.AccountController) AccountHandler {

	return &accountHandler{
		accountController: accountController,
	}
}
func (h *accountHandler) GetRootHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Printf("%s GET request\n", ctx.Value(constants.ADDRESS))
	io.WriteString(w, "EBANX Challenge Root\n")
}

func (h *accountHandler) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("%s/balance GET request\n", ctx.Value(constants.ADDRESS))

	accountID := r.URL.Query().Get("account_id")


	amount, err := h.accountController.GetBalanceController(accountID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Get Balance Error: %s account_id: %s", err.Error(), accountID)
		res := fmt.Sprintf("%d %d", http.StatusNotFound, 0)
		io.WriteString(w, res)
		return
	}
	res := fmt.Sprintf("%d %s", http.StatusOK, amount)
	io.WriteString(w, res)

}

func (h *accountHandler) PostEventHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Printf("%s/event POST request\n", ctx.Value(constants.ADDRESS))
	if req.Body == nil {
		w.WriteHeader(http.StatusNotFound)
		res := fmt.Sprintf("%d %d", http.StatusNotFound, 0)
		log.Printf("Account Event Error: %s", errors.ErrMissingBody.Error())
		io.WriteString(w, res)
		return
	}
	decoder := json.NewDecoder(req.Body)

	var event model.Event
	err := decoder.Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res := fmt.Sprintf("%d %d", http.StatusNotFound, 0)
		log.Printf("Account Event Error: %s", err.Error())
		io.WriteString(w, res)
		return
	}

	res := ""

	switch event.Type {
	case "deposit":
		res, err = h.accountController.PostDepositEventController(*event.Destination, event.Amount)
	case "withdraw":
		res, err = h.accountController.PostWithdrawEventController(*event.Origin, event.Amount)
	case "transfer":
		res, err = h.accountController.PostTransferEventController(*event.Origin, *event.Destination, event.Amount)
	default:
		err = errors.ErrInvalidEventType
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res := fmt.Sprintf("%d %d", http.StatusNotFound, 0)
		log.Printf("Account Event Error: %s", err.Error())
		io.WriteString(w, res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf("%d %s", http.StatusCreated, res)
	io.WriteString(w, response)

}

func (h *accountHandler) PostDeleteHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Printf("%s/delete POST request\n", ctx.Value(constants.ADDRESS))

	h.accountController.DeleteAllAccountsController()
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, strconv.Itoa(http.StatusOK))
}

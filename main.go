package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	// "os"

	"github.com/nkuros/ebanxchallenge/constants"
	"github.com/nkuros/ebanxchallenge/handler"
	"github.com/nkuros/ebanxchallenge/service"
	"github.com/nkuros/ebanxchallenge/controller"
)


func main() {

	ctx := context.Background()

	accountService := service.NewAccountService()
	accountController := controller.NewAccountController(accountService)
	accountHandler := handler.NewAccountHandler(accountController)


	handlers := http.NewServeMux()
	handlers.HandleFunc("/", accountHandler.GetRootHandler)
	handlers.HandleFunc("/balance", accountHandler.GetBalanceHandler)
	handlers.HandleFunc("/event", accountHandler.PostEventHandler)
	handlers.HandleFunc("/delete", accountHandler.PostDeleteHandler)

	server := &http.Server{
		Addr:    constants.PORT,
		Handler: handlers,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, constants.ADDRESS, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	} else if err != nil {
		log.Fatal(err)
		// os.Exit(1)
	}
}

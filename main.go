package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/nkuros/ebanxchallenge/controller"
	"github.com/nkuros/ebanxchallenge/handler"
	"github.com/nkuros/ebanxchallenge/service"
	"github.com/spf13/viper"
)


func main() {

	ctx := context.Background()

	accountService := service.NewAccountService()
	accountController := controller.NewAccountController(accountService)
	accountHandler := handler.NewAccountHandler(accountController)
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	address := viper.GetString("ADDRESS")
	port := viper.GetString("PORT")
	
	handlers := http.NewServeMux()
	handlers.HandleFunc("/", accountHandler.GetRootHandler)
	handlers.HandleFunc("/balance", accountHandler.GetBalanceHandler)
	handlers.HandleFunc("/event", accountHandler.PostEventHandler)
	handlers.HandleFunc("/delete", accountHandler.PostDeleteHandler)

	server := &http.Server{
		Addr:    address + ":" + port,
		Handler: handlers,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, address, l.Addr().String())
			return ctx
		},
	}

	err = server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	} else if err != nil {
		log.Fatal(err)
	}
}

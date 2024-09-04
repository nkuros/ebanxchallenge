package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/nkuros/ebanxchallenge/constants"
	"github.com/nkuros/ebanxchallenge/handler"
)



func main() {

	ctx := context.Background()

	handlers := http.NewServeMux()
	handlers.HandleFunc("/", handler.GetRoot)

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
		os.Exit(1)
	}
}
package handler

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/nkuros/ebanxchallenge/constants"
)



func GetRootHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("%s GET request\n", ctx.Value(constants.ADDRESS))
	io.WriteString(w, "EBANX Challenge Root\n")
}

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("%s/balance GET request\n", ctx.Value(constants.ADDRESS))
	hasAccIdQuery := r.URL.Query().Has("id")
	if hasAccIdQuery == true {
		_, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			
			io.WriteString(w, "404 invalid id format")
			return
		}
		io.WriteString(w, "200")
		// TODO: Add Getbalance Controller
	} else {
		io.WriteString(w, "404 0")
	}
}

func PostEventHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	eventType := r.PostFormValue("type")
	if eventType == ""{
		io.WriteString(w, "404 0")
		return
	}
	_, err := strconv.Atoi(r.PostFormValue("origin"))
	if err != nil {
			
		io.WriteString(w, "404 invalid id format")
		return
	}
	_, err = strconv.Atoi(r.PostFormValue("amount"))
	if err != nil {
			
		io.WriteString(w, "404 invalid amount format")
		return
	}
	// destination := r.PostFormValue("destination")

	// TODO: Add PostEvent Controller
	log.Printf("%s/event POST request\n", ctx.Value(constants.ADDRESS))
	io.WriteString(w, "200")
}

func PostDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// TODO: Add PostDelete Controller
	log.Printf("%s/delete POST request\n", ctx.Value(constants.ADDRESS))
	io.WriteString(w, "200")
}

package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/nkuros/ebanxchallenge/constants"
)



func GetRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("%s GET request\n", ctx.Value(constants.ADDRESS))
	io.WriteString(w, "EBANX Challenge Root\n")
}
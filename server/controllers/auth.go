package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorvk/angular-go-auth/server/util"
)

const wrongPlaceMessage = "bad request"
const wrongMethodMessage = "method not allowed"

func Login(w http.ResponseWriter, r *http.Request) {
	handleResponse(w, r, http.MethodPost, func() {
		d, err := io.ReadAll(r.Body)
		util.HandleIfHttpError(err, w, wrongPlaceMessage, http.StatusBadRequest)
		fmt.Fprintf(w, "Please login to your account %s", d)
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	handleResponse(w, r, http.MethodPost, func() {
		d, err := io.ReadAll(r.Body)
		util.HandleIfHttpError(err, w, wrongPlaceMessage, http.StatusBadRequest)
		fmt.Fprintf(w, "Please create your account %s", d)
	})
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
}

func handleResponse(w http.ResponseWriter, r *http.Request, methodType string, handler func()) {
	if r.Method == methodType {
		enableCORS(w)
		handler()
		return
	}
	http.Error(w, wrongMethodMessage, http.StatusMethodNotAllowed)
}

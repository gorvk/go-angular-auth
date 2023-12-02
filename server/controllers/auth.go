package controllers

import (
	"fmt"
	"io"
	"net/http"
)

const wrongPlaceMessage = "bad request"
const wrongMethodMessage = "method not allowed"

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		d, err := io.ReadAll(r.Body)
		handleIfHttpError(err, w, wrongPlaceMessage, http.StatusBadRequest)
		fmt.Fprintf(w, "Please login to your account %s", d)
		return
	}
	http.Error(w, wrongMethodMessage, http.StatusMethodNotAllowed)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		d, err := io.ReadAll(r.Body)
		handleIfHttpError(err, w, wrongPlaceMessage, http.StatusBadRequest)
		fmt.Fprintf(w, "Please create your account %s", d)
		return
	}
	http.Error(w, wrongMethodMessage, http.StatusMethodNotAllowed)
}

func handleIfHttpError(err error, w http.ResponseWriter, errorMsg string, code int) {
	if err != nil {
		http.Error(w, errorMsg, code)
	}
}

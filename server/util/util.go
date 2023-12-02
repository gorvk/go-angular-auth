package util

import "net/http"

func HandleIfHttpError(err error, w http.ResponseWriter, errorMsg string, code int) {
	if err != nil {
		http.Error(w, errorMsg, code)
	}
}

package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorvk/angular-go-auth/server/models"
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
		var t models.User = models.User{}
		err := json.NewDecoder(r.Body).Decode(&t)
		result, err := models.InsertUser(t)
		rowsAffected, _ := result.RowsAffected()
		d, err := json.Marshal(
			map[string]interface{}{
				"success":                 true,
				"number of rows affected": rowsAffected,
			},
		)
		util.HandleIfHttpError(err, w, wrongPlaceMessage, http.StatusBadRequest)
		fmt.Fprintf(w, "%v", d)
	})
}

func handleResponse(w http.ResponseWriter, r *http.Request, methodType string, handler func()) {
	if r.Method == methodType {
		handler()
		return
	}
	http.Error(w, wrongMethodMessage, http.StatusMethodNotAllowed)
}

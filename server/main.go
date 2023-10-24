package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorvk/angular-go-auth/server/auth"
	"github.com/gorvk/angular-go-auth/server/types"
)

func main() {
	server := mux.NewRouter()

	enableCORS := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			next.ServeHTTP(w, r)
		})
	}

	server.Use(enableCORS)
	server.HandleFunc("/register", auth.Register)
	server.NotFoundHandler = enableCORS(http.HandlerFunc(NotFoundHandler))
	server.HandleFunc("/", auth.Login)
	fmt.Println("server running at http://localhost:3333")
	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("error while opening the server")
	}
}

func NotFoundHandler(reponseWriter http.ResponseWriter, _ *http.Request) {
	s, _ := json.Marshal(
		types.Response{
			Success:    false,
			StatusCode: http.StatusNotFound,
			Data: []auth.User{
				{
					Id:        0,
					Firstname: "xxx",
					Lastname:  "xxx",
				},
			},
		},
	)
	http.Error(reponseWriter, string(s), http.StatusNotFound)

}

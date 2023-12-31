package routes

import (
	"net/http"

	"github.com/gorvk/angular-go-auth/server/controllers"
)

func Setup() {
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
}

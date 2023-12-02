package routes

import (
	"net/http"

	"github.com/gorvk/angular-go-auth/server/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
}

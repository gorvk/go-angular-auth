package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorvk/angular-go-auth/server/types"
)

type User struct {
	Id        int
	Firstname string
	Lastname  string
}

func Login(reponseWriter http.ResponseWriter, requestReader *http.Request) {
	if requestReader.Method == http.MethodGet {
		user := []User{
			{
				Id:        1,
				Firstname: "Gourav",
				Lastname:  "Kolhatkar",
			},
			{
				Id:        2,
				Firstname: "Nayan",
				Lastname:  "Kolhatkar",
			},
			{
				Id:        3,
				Firstname: "Yukta",
				Lastname:  "Kolhatkar",
			},
		}

		json.NewEncoder(reponseWriter).Encode(
			types.Response{
				Success:    true,
				StatusCode: http.StatusOK,
				Data:       user,
			},
		)
	}
}

func Register(responseWriter http.ResponseWriter, requestReader *http.Request) {
	if requestReader.Method == http.MethodGet {
		user := []User{
			{
				Id:        1,
				Firstname: "Pranjali",
				Lastname:  "Bramhankar",
			},
			{
				Id:        2,
				Firstname: "Sagar",
				Lastname:  "Biswas",
			},
		}
		json.NewEncoder(responseWriter).Encode(
			types.Response{
				Success:    true,
				StatusCode: http.StatusOK,
				Data:       user,
			},
		)
	}
}

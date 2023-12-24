package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/angular-go-auth/server/database"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func InsertUser(user User) (sql.Result, error) {
	var db *sql.DB = database.GetDbInstance()
	var query = fmt.Sprintf("insert into userdetails values('%v', %v, '%v')", user.Username, user.Age, user.Password)
	return db.Exec(query)
}

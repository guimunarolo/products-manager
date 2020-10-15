package models

import (
    "time"
	"github.com/guimunarolo/products-manage/calculator-service/config"
)

// User struct from table `users`
type User struct {
    tableName struct{} `pg:"users,alias:u"`

    ID          string
    FirstName   string
    LastName    string
    DateOfBirth  time.Time
}


func init() {
	config.Connect()
	db = config.GetDB()
}

// GetUserByID returns a *User selected by given ID
func GetUserByID(userID string) (*User){
    user := &User{ID: userID}
    err := db.Model(user).WherePK().Select()
    if err != nil {
        panic(err)
    }
    
    return user
}

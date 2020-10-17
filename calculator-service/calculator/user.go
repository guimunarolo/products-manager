package calculator

import "time"

type User struct {
    tableName struct{} `pg:"users,alias:u"`

    ID          string
    FirstName   string
    LastName    string
    DateOfBirth time.Time
}

type UserRespository interface {
	GetUser(id string) (*User, error)
}

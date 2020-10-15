package config

import (
	"os"

	"github.com/go-pg/pg"
)

var db *pg.DB

func Connect() {
	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil{
		panic(err)
	}

	db = pg.Connect(opt)
}

func GetDB() *pg.DB {
	return db
}

package models

import (
	"github.com/go-pg/pg"
	"github.com/guimunarolo/products-manage/calculator-service/config"
)

var db *pg.DB

// Product struct from table `products`
type Product struct {
    tableName struct{} `pg:"products,alias:u"`

    ID           string
	Title	     string
    Description  string
    PriceInCents int32
}


func init() {
	config.Connect()
	db = config.GetDB()
}

// GetProductByID returns a *User selected by given ID
func GetProductByID(productID string) (*Product){
    product := &Product{ID: productID}
    err := db.Model(product).WherePK().Select()
    if err != nil {
        panic(err)
    }
    
    return product
}

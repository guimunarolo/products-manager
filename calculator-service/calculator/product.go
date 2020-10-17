package calculator

type Product struct {
    tableName struct{} `pg:"products,alias:u"`

    ID           string
	Title	     string
    Description  string
    PriceInCents int32
}

type ProductRepository interface {
	GetProduct(id string) (*Product, error)
}

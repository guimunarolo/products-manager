package calculator

import (
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/hashicorp/go-hclog"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *pg.DB
	logger hclog.Logger
}

func NewUserRepo(db *pg.DB, logger hclog.Logger) UserRespository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func NewProductRepo(db *pg.DB, logger hclog.Logger) ProductRepository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (repo *repo) GetUser(id string) (*User, error) {
	user := &User{ID: id}
    err := repo.db.Model(user).WherePK().Select()
    if err != nil {
		return nil, RepoErr
	}

    return user, nil
}

func (repo *repo) GetProduct(id string) (*Product, error) {
	product := &Product{ID: id}
    err := repo.db.Model(product).WherePK().Select()
    if err != nil {
		return nil, RepoErr
	}

    return product, nil
}

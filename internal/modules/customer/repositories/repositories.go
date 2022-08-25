package repositories

import (
	"database/sql"
)

type Repositories interface {
	RepositoryCustomerQuery
}

type RepositoriesImpl struct {
	// inject db impl to RepositoriesImpl event the db is being used by the child struct impl
	DB *sql.DB
	*RepositoryCustomerQueryImpl
}

func NewCustomerRepository(
	db *sql.DB,
) *RepositoriesImpl {
	return &RepositoriesImpl{
		DB: db,
		RepositoryCustomerQueryImpl: &RepositoryCustomerQueryImpl{
			DB: db,
		},
	}
}

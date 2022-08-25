package repositories

import (
	"database/sql"
)

type Repositories interface {
	RepositoryCompanyPersonQuery
}

type RepositoriesImpl struct {
	// inject db impl to RepositoriesImpl event the db is being used by the child struct impl
	db *sql.DB
	*RepositoryCompanyPersonQueryImpl
}

func NewCompanyPersonRepository(
	db *sql.DB,
) *RepositoriesImpl {
	return &RepositoriesImpl{
		db: db,
		RepositoryCompanyPersonQueryImpl: &RepositoryCompanyPersonQueryImpl{
			DB: db,
		},
	}
}

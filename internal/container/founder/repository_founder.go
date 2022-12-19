package founder

import (
	"kueku/internal/container"
	"kueku/internal/domain/cake"
	"kueku/internal/persistence"
	repo "kueku/internal/persistence/repository/sqlx"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	DB    *persistence.DB
	query container.Queries

	cakeRepo cake.Repository
}

// CakeRepository .
func (r *repository) CakeRepository() cake.Repository {
	if r.cakeRepo == nil {
		r.cakeRepo = &repo.Cake{DB: r.DB, Query: r.query.CakeQueries()}
	}
	return r.cakeRepo
}

// NewRepository .
func NewRepository(DB *sqlx.DB, query container.Queries) container.Repository {
	return &repository{DB: &persistence.DB{DB: DB}, query: query}
}

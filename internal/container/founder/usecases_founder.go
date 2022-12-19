package founder

import (
	"kueku/internal/container"
	"kueku/internal/usecase"
)

type usecases struct {
	repo  container.Repository
	query container.Queries

	cakeUsecases usecase.CakeUsecase
}

// CakeUsecase .
func (c *usecases) CakeUsecase() usecase.CakeUsecase {
	if c.cakeUsecases == nil {
		c.cakeUsecases = usecase.NewCake(c.repo.CakeRepository())
	}
	return c.cakeUsecases

}

// NewUsecases .
func NewUsecases(repo container.Repository, query container.Queries) container.Usecases {
	return &usecases{repo: repo, query: query}
}

package usecase

import (
	"context"
	"kueku/internal/usecase/model"
	"kueku/internal/domain/cake"
)

type Cake struct {
	repo cake.Repository
}

// Usecase .
type CakeUsecase interface {
	List(ctx context.Context) (cake.Cakes, error)
	Detail(ctx context.Context, id int) (*cake.Cake, error)
	Create(ctx context.Context, command *model.CreateCakeCommand) (*cake.Cake, error)
	Update(ctx context.Context, command *model.UpdateCakeCommand) (*cake.Cake, error)
	Delete(ctx context.Context, id int) error
}

// NewCake .
func NewCake(repo cake.Repository) CakeUsecase {
	return &Cake{repo: repo}
}

func (c *Cake) List(ctx context.Context) (cake.Cakes, error) {
	return c.repo.Fetch(ctx)
}

func (c *Cake) Detail(ctx context.Context, id int) (*cake.Cake, error) {
	return c.repo.GetByID(ctx, id)

}

func (c *Cake) Create(ctx context.Context, command *model.CreateCakeCommand) (*cake.Cake, error) {
	data := command.ToCake()
	lid, err := c.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return c.repo.GetByID(ctx, lid)
}

func (c *Cake) Update(ctx context.Context, command *model.UpdateCakeCommand) (*cake.Cake, error) {
	data := command.ToCake()
	if err := c.repo.Update(ctx, data); err != nil {
		return nil, err
	}

	return c.repo.GetByID(ctx, data.ID)
}

func (c *Cake) Delete(ctx context.Context, id int) error {
	return c.repo.Delete(ctx, id)

}

package cake

import "context"

// Repository . 
type Repository interface {
	Fetch(ctx context.Context) (Cakes, error)
	GetByID(ctx context.Context, id int) (*Cake, error)
	Create(ctx context.Context, data *Cake) (int, error)
	Update(ctx context.Context, data *Cake) error
	Delete(ctx context.Context, id int) error
}

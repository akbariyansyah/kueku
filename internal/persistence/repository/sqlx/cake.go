package sqlx

import (
	"context"
	"database/sql"
	"kueku/internal/domain/cake"
	"kueku/internal/persistence"
	"kueku/internal/persistence/queries"
	"kueku/internal/persistence/repository"
)

// Cake .
type Cake struct {
	DB    persistence.Sqlx
	Query queries.Cake
}

// Fetch .
func (r *Cake) Fetch(ctx context.Context) (cake.Cakes, error) {
	query, args, err := r.Query.Fetch()
	if err != nil {
		return nil, repository.NewErrQuery(err)
	}

	rows, err := r.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, repository.NewErrDatabase(err)
	}

	var cakes cake.Cakes

	for rows.Next() {
		c := new(cake.Cake)
		if err := rows.Scan(
			&c.ID,
			&c.Title,
			&c.Description,
			&c.Image,
			&c.Rating,
			&c.CreatedAt,
			&c.UpdatedAt,
		); err != nil {
			return nil, repository.NewErrDatabase(err)
		}
		cakes = append(cakes, c)
	}
	return cakes, nil
}

// GetByID .
func (r *Cake) GetByID(ctx context.Context, id int) (*cake.Cake, error) {
	query, args, err := r.Query.GetByID(id)
	if err != nil {
		return nil, repository.NewErrQuery(err)
	}

	p := new(cake.Cake)
	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(
		&p.ID,
		&p.Title,
		&p.Description,
		&p.Image,
		&p.Rating,
		&p.CreatedAt,
		&p.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, cake.ErrNotFound
		}
		return nil, repository.NewErrDatabase(err)

	}

	return p, nil
}

// Create .
func (r *Cake) Create(ctx context.Context, data *cake.Cake) error {
	query, args, err := r.Query.Create(data)
	if err != nil {
		return repository.NewErrQuery(err)
	}
	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		return repository.NewErrDatabase(err)
	}
	return nil
}

// Update .
func (r *Cake) Update(ctx context.Context, data *cake.Cake) error {
	query, args, err := r.Query.Update(data)
	if err != nil {
		return repository.NewErrQuery(err)
	}
	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		return repository.NewErrDatabase(err)

	}
	return nil
}

// Update .
func (r *Cake) Delete(ctx context.Context, id int) error {
	query, args, err := r.Query.Delete(id)
	if err != nil {
		return repository.NewErrQuery(err)
	}

	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		return repository.NewErrDatabase(err)

	}
	return nil
}

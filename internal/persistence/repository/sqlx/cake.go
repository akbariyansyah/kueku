package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"kueku/internal/domain/cake"
	"kueku/internal/persistence/queries"
	"kueku/internal/persistence/repository"

	"github.com/jmoiron/sqlx"
)

// Cake .
type Cake struct {
	DB    *sqlx.DB
	Query queries.Cake
}

// Fetch .
func (r *Cake) Fetch(ctx context.Context) (cake.Cakes, error) {
	query, args, err := r.Query.Fetch()
	if err != nil {
		return nil, fmt.Errorf("%q : %e", repository.ErrQuery, err)
	}

	rows, err := r.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%q : %e", repository.ErrDatabase, err)

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
			return nil, fmt.Errorf("%q : %e", repository.ErrDatabase, err)
		}
		cakes = append(cakes, c)
	}
	return cakes, nil
}

// GetByID .
func (r *Cake) GetByID(ctx context.Context, id int) (*cake.Cake, error) {
	query, args, err := r.Query.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%q : %e", repository.ErrQuery, err)
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
		return nil, fmt.Errorf("%q : %e", repository.ErrDatabase, err)

	}

	return p, nil
}

// Create .
func (r *Cake) Create(ctx context.Context, data *cake.Cake) error {
	query, args, err := r.Query.Create(data)
	if err != nil {
		return fmt.Errorf("%q : %e", repository.ErrQuery, err)
	}
	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("%q : %e", repository.ErrDatabase, err)
	}
	return nil
}

// Update .
func (r *Cake) Update(ctx context.Context, data *cake.Cake) error {
	query, args, err := r.Query.Update(data)
	if err != nil {
		return fmt.Errorf("%q : %e", repository.ErrQuery, err)
	}
	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("%q : %e", repository.ErrDatabase, err)

	}
	return nil
}

// Update .
func (r *Cake) Delete(ctx context.Context, id int) error {
	query, args, err := r.Query.Delete(id)
	if err != nil {
		return fmt.Errorf("%q : %e", repository.ErrQuery, err)
	}
	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("%q : %e", repository.ErrDatabase, err)

	}
	return nil
}

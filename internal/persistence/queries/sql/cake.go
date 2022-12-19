package sql

import (
	"kueku/internal/domain/cake"

	sq "github.com/Masterminds/squirrel"
)

const (
	cakeTable = "cakes"
)

type Cake struct {
	Sql sq.StatementBuilderType
}

// Fetch .
func (c *Cake) Fetch() (string, []interface{}, error) {
	sql := c.Sql.Select(
		"id",
		"title",
		"description",
		"image",
		"rating",
		"created_at",
		"updated_at",
	).From(cakeTable)

	return sql.ToSql()
}

// GetByID .
func (c *Cake) GetByID(id int) (string, []interface{}, error) {
	sql := c.Sql.Select(
		"id",
		"title",
		"description",
		"image",
		"rating",
		"created_at",
		"updated_at",
	).From(cakeTable).Where(
		sq.Eq{"id": id},
	)
	return sql.ToSql()
}

// Create .
func (c *Cake) Create(data *cake.Cake) (string, []interface{}, error) {
	sql := c.Sql.Insert(cakeTable).Columns(
		"id",
		"title",
		"description",
		"image",
		"rating",
		"created_at",
		"updated_at",
	).Values(
		data.ID,
		data.Title,
		data.Description,
		data.Image,
		data.Rating,
		data.CreatedAt,
		data.UpdatedAt,
	)
	return sql.ToSql()
}

// Update .
func (c *Cake) Update(data *cake.Cake) (string, []interface{}, error) {
	sql := c.Sql.Update(cakeTable).
		Set("title", data.Title).
		Set("description", data.Description).
		Set("image", data.Image).
		Set("rating", data.Rating).
		Set("updated_at", data.UpdatedAt).
		Where(sq.Eq{"id": data.ID})

	return sql.ToSql()
}

// Update .
func (c *Cake) Delete(id int) (string, []interface{}, error) {
	sql := c.Sql.Delete(cakeTable).Where(sq.Eq{"id": id})
	return sql.ToSql()
}

package queries

import (
	"kueku/internal/domain/cake"
)

// Cake .
type Cake interface {
	Fetch() (string, []interface{}, error)
	GetByID(id int) (string, []interface{}, error)
	Create(data *cake.Cake) (string, []interface{}, error)
	Update(data *cake.Cake) (string, []interface{}, error)
	Delete(id int) (string, []interface{}, error)
}

package container

import (
	"kueku/internal/persistence/queries"

)

// Queries .
type Queries interface {
	CakeQueries() queries.Cake
}

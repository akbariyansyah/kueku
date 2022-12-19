package founder

import (
	"kueku/internal/container"
	"kueku/internal/persistence/queries"

	sq "github.com/Masterminds/squirrel"
	"kueku/internal/persistence/queries/sql"
)

type appQuery struct {
	format sq.PlaceholderFormat

	cakeQueries queries.Cake
}

// CakeQueries .
func (q *appQuery) CakeQueries() queries.Cake {
	if q.cakeQueries == nil {
		q.cakeQueries = &sql.Cake{Sql: sq.StatementBuilder.PlaceholderFormat(q.format)}
	}
	return q.cakeQueries
}

// NewQueries .
func NewQueries(format sq.PlaceholderFormat) container.Queries {
	return &appQuery{format: format}
}

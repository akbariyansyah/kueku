package sql_test

import (
	"kueku/internal/domain/cake"
	"kueku/internal/persistence/queries"
	"kueku/internal/persistence/queries/sql"
	"testing"

	sq "github.com/Masterminds/squirrel"

	"github.com/stretchr/testify/suite"
)

type cakeTestSuite struct {
	suite.Suite
	querier queries.Cake
}

func TestPendingPostTransaction(t *testing.T) {
	suite.Run(t, new(cakeTestSuite))
}

func (t *cakeTestSuite) SetupTest() {
	t.querier = &sql.Cake{Sql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func (t *cakeTestSuite) TestFetch_NoError() {
	query, args, err := t.querier.Fetch()

	t.NoError(err)
	t.Len(args, 0)
	t.Contains(query, "SELECT")
}

func (t *cakeTestSuite) TestGetByID_NoError() {
	mockID := 5

	query, args, err := t.querier.GetByID(mockID)

	t.NoError(err)
	t.Len(args, 1)
	t.Contains(query, "SELECT")
}

func (t *cakeTestSuite) TestCreate_NoError() {
	mockData := new(cake.Cake)

	query, args, err := t.querier.Create(mockData)

	t.NoError(err)
	t.Len(args, 7)
	t.Contains(query, "INSERT")
}

func (t *cakeTestSuite) TestUpdate_NoError() {
	mockData := new(cake.Cake)

	query, args, err := t.querier.Update(mockData)

	t.NoError(err)
	t.Len(args, 6)
	t.Contains(query, "UPDATE")
}

func (t *cakeTestSuite) TestDelete_NoError() {
	mockID := 5

	query, args, err := t.querier.Delete(mockID)

	t.NoError(err)
	t.Len(args, 1)
	t.Contains(query, "DELETE")
}

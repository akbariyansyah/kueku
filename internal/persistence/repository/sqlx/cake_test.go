package sqlx_test

import (
	"context"
	"database/sql"
	"errors"
	"kueku/internal/domain/cake"
	"kueku/internal/persistence"
	pmocks "kueku/internal/persistence/mocks"
	"kueku/internal/persistence/queries"
	qmocks "kueku/internal/persistence/queries/mocks"
	qsql "kueku/internal/persistence/queries/sql"
	"kueku/internal/persistence/repository"
	reposqlx "kueku/internal/persistence/repository/sqlx"

	sq "github.com/Masterminds/squirrel"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type cakeTestSuite struct {
	suite.Suite
	ctx         context.Context
	mockCtrl    *gomock.Controller
	mockSQL     sqlmock.Sqlmock
	mockQuerier *qmocks.MockCake
	mockDB      *pmocks.MockSqlx
	querier     queries.Cake
	db          *persistence.DB
	repo        cake.Repository
}

func TestCakeSuite(t *testing.T) {
	suite.Run(t, new(cakeTestSuite))
}

func (t *cakeTestSuite) SetupTest() {
	var mockSQL *sql.DB

	t.ctx = context.Background()
	t.mockCtrl = gomock.NewController(t.T())
	mockSQL, t.mockSQL, _ = sqlmock.New()
	mockSqlxDB := sqlx.NewDb(mockSQL, "sqlmock")
	t.mockQuerier = qmocks.NewMockCake(t.mockCtrl)
	t.mockDB = pmocks.NewMockSqlx(t.mockCtrl)
	t.querier = &qsql.Cake{Sql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
	t.db = &persistence.DB{DB: mockSqlxDB}
	t.repo = &reposqlx.Cake{DB: t.mockDB, Query: t.querier}
}

func (t *cakeTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
}

func (t *cakeTestSuite) TestFetch_QuerierError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.mockQuerier}

	mockErr := errors.New("unexpected")

	t.mockQuerier.EXPECT().Fetch().Return("", nil, mockErr)

	res, err := t.repo.Fetch(t.ctx)

	t.Nil(res)
	t.Equal(err, repository.NewErrQuery(mockErr))
}

func (t *cakeTestSuite) TestFetch_QueryxContextError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}

	mockErr := errors.New("unexpected")

	t.mockSQL.ExpectQuery("^SELECT (.+)").WillReturnError(mockErr)

	res, err := t.repo.Fetch(t.ctx)

	t.Nil(res)
	t.Equal(err, repository.NewErrDatabase(mockErr))
}

func (t *cakeTestSuite) TestFetch_ScanError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}
	mockErr := errors.New("sql: expected 1 destination arguments in Scan, not 7")

	row := &cake.Cake{}
	rows := sqlmock.NewRows([]string{
		"id",
	})
	rows.AddRow(row.ID).RowError(1, mockErr)

	t.mockSQL.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)

	res, err := t.repo.Fetch(t.ctx)

	t.Nil(res)
	t.Error(err)
	t.NoError(t.mockSQL.ExpectationsWereMet())
}

func (t *cakeTestSuite) TestFetch_NoError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}

	mockRows := cake.Cakes{
		&cake.Cake{ID: 1},
	}
	row := mockRows[0]

	rows := sqlmock.NewRows([]string{
		"id",
		"title",
		"description",
		"image",
		"rating",
		"created_at",
		"updated_at",
	})
	rows.AddRow(
		row.ID,
		row.Title,
		row.Description,
		row.Image,
		row.Rating,
		row.CreatedAt,
		row.UpdatedAt,
	)

	t.mockSQL.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)

	res, err := t.repo.Fetch(t.ctx)

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Equal(mockRows, res)
	t.NoError(err)
}

func (t *cakeTestSuite) TestGetByID_QuerierError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.mockQuerier}

	mockErr := errors.New("unexpected")
	mockID := 10
	t.mockQuerier.EXPECT().GetByID(mockID).Return("", nil, mockErr)

	res, err := t.repo.GetByID(t.ctx, mockID)

	t.Nil(res)
	t.Equal(err, repository.NewErrQuery(mockErr))
}

func (t *cakeTestSuite) TestGetByID_QueryxContextError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}

	mockErr := errors.New("unexpected")
	mockID := 10

	t.mockSQL.ExpectQuery("^SELECT (.+)").WillReturnError(mockErr)

	res, err := t.repo.GetByID(t.ctx, mockID)

	t.Nil(res)
	t.Equal(err, repository.NewErrDatabase(mockErr))
}

func (t *cakeTestSuite) TestGetByID_ErrDataNotFound() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}
	mockID := 10

	t.mockSQL.ExpectQuery("^SELECT (.+)").WillReturnError(sql.ErrNoRows)

	res, err := t.repo.GetByID(t.ctx, mockID)

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Nil(res)
	t.Equal(cake.ErrNotFound, err)
}

func (t *cakeTestSuite) TestGetByID_NoError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}

	mockRows := cake.Cakes{
		&cake.Cake{ID: 1},
	}
	row := mockRows[0]
	mockID := 10

	rows := sqlmock.NewRows([]string{
		"id",
		"title",
		"description",
		"image",
		"rating",
		"created_at",
		"updated_at",
	})
	rows.AddRow(
		row.ID,
		row.Title,
		row.Description,
		row.Image,
		row.Rating,
		row.CreatedAt,
		row.UpdatedAt,
	)

	t.mockSQL.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)

	res, err := t.repo.GetByID(t.ctx, mockID)

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.NotNil(res)
	t.NoError(err)
}

func (t *cakeTestSuite) TestDelete_QueryerError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.mockQuerier}

	mockID := 10
	mockErr := errors.New("unexpected")

	t.mockQuerier.EXPECT().Delete(mockID).Return("", nil, mockErr)

	err := t.repo.Delete(t.ctx, mockID)
	t.Equal(err, repository.NewErrQuery(mockErr))
}

func (t *cakeTestSuite) TestDelete_ExecContextError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}
	mockID := 10
	mockErr := errors.New("unexpected")

	t.mockSQL.ExpectExec("^DELETE (.+)").
		WithArgs(mockID).
		WillReturnError(mockErr)

	err := t.repo.Delete(t.ctx, mockID)

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Equal(err, repository.NewErrDatabase(mockErr))

}

func (t *cakeTestSuite) TestDelete_NoError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}
	mockID := 10

	t.mockSQL.ExpectExec("^DELETE (.+)").
		WithArgs(mockID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := t.repo.Delete(t.ctx, mockID)

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.NoError(err)
}

func (t *cakeTestSuite) TestCreate_QuerierError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.mockQuerier}

	mockData := new(cake.Cake)
	mockErr := errors.New("unexpected")

	t.mockQuerier.EXPECT().Create(mockData).Return("", nil, mockErr)

	err := t.repo.Create(t.ctx, mockData)

	t.Equal(err, repository.NewErrQuery(mockErr))
}

func (t *cakeTestSuite) TestCreate_ExecContextError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}

	mockData := new(cake.Cake)

	mockErr := errors.New("unexpected")

	t.mockSQL.ExpectExec("^INSERT (.+)").WithArgs(
		mockData.ID,
		mockData.Title,
		mockData.Description,
		mockData.Image,
		mockData.Rating,
		mockData.CreatedAt,
		mockData.UpdatedAt,
	).WillReturnError(mockErr)

	err := t.repo.Create(t.ctx, mockData)

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Equal(err, repository.NewErrDatabase(mockErr))
}

func (t *cakeTestSuite) TestCreate_NoError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}

	mockData := &cake.Cake{}

	t.mockSQL.ExpectExec("^INSERT (.+)").WithArgs(
		mockData.ID,
		mockData.Title,
		mockData.Description,
		mockData.Image,
		mockData.Rating,
		mockData.CreatedAt,
		mockData.UpdatedAt,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err := t.repo.Create(t.ctx, mockData)

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.NoError(err)
}

func (t *cakeTestSuite) TestUpdate_QuerierError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.mockQuerier}

	mockData := new(cake.Cake)
	mockErr := errors.New("unexpected")

	t.mockQuerier.EXPECT().Update(mockData).Return("", nil, mockErr)

	err := t.repo.Update(t.ctx, mockData)

	t.Equal(err, repository.NewErrQuery(mockErr))
}

func (t *cakeTestSuite) TestUpdate_ExecContextError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}

	mockData := &cake.Cake{}
	mockErr := errors.New("unexpected")

	t.mockSQL.ExpectExec("^UPDATE (.+)").WithArgs(
		mockData.Title,
		mockData.Description,
		mockData.Image,
		mockData.Rating,
		mockData.UpdatedAt,
		mockData.ID,
	).WillReturnError(mockErr)

	err := t.repo.Update(t.ctx, mockData)

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Equal(err, repository.NewErrDatabase(mockErr))
}

func (t *cakeTestSuite) TestUpdate_NoError() {
	t.repo = &reposqlx.Cake{DB: t.db, Query: t.querier}

	mockData := &cake.Cake{}

	t.mockSQL.ExpectExec("^UPDATE (.+)").WithArgs(
		mockData.Title,
		mockData.Description,
		mockData.Image,
		mockData.Rating,
		mockData.UpdatedAt,
		mockData.ID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err := t.repo.Update(t.ctx, mockData)

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.NoError(err)
}



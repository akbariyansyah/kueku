package persistence_test

import (
	"context"
	"database/sql"
	"errors"
	"kueku/internal/persistence"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type dbTestSuite struct {
	suite.Suite
	ctx        context.Context
	mockCtrl   *gomock.Controller
	mockDB     *sql.DB
	mockSqlxDB *sqlx.DB
	mockSQL    sqlmock.Sqlmock
	db         persistence.DB
}

func TestDBTestSuite(t *testing.T) {
	suite.Run(t, new(dbTestSuite))
}

func (t *dbTestSuite) SetupTest() {
	t.ctx = context.Background()
	t.mockCtrl = gomock.NewController(t.T())
	t.mockDB, t.mockSQL, _ = sqlmock.New()
	t.mockSqlxDB = sqlx.NewDb(t.mockDB, "sqlmock")
	t.db = persistence.DB{t.mockSqlxDB}
	t.db.Startup()
}

func (t *dbTestSuite) TearDownTest() {
	t.db.Shutdown()
	t.mockCtrl.Finish()
}

func (t *dbTestSuite) TestPing_NoError() {
	err := t.db.Ping(t.ctx)
	t.NoError(err)
}

func (t *dbTestSuite) TestTx_BeginError() {
	mockErr := errors.New("unexpected")

	t.mockSQL.ExpectBegin().WillReturnError(mockErr)

	err := t.db.Tx(t.ctx, func(ctx context.Context, err chan error) {
		err <- nil
		return
	})

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Error(err)
}

func (t *dbTestSuite) TestTx_CommitNoError() {
	t.mockSQL.ExpectBegin()
	t.mockSQL.ExpectCommit()

	err := t.db.Tx(t.ctx, func(ctx context.Context, err chan error) {
		err <- nil
		return
	})

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.NoError(err)
}

func (t *dbTestSuite) TestNestedTx_Commit_NoError() {
	t.mockSQL.ExpectBegin()
	t.mockSQL.ExpectCommit()

	err := t.db.Tx(t.ctx, func(ctx context.Context, err chan error) {
		err <- t.db.Tx(ctx, func(ctx context.Context, err chan error) {
			err <- t.db.Tx(ctx, func(ctx context.Context, err chan error) {
				err <- nil
				return
			})
			return
		})
		return
	})

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.NoError(err)
}

func (t *dbTestSuite) TestTx_Rollback_Error() {
	mockErr := errors.New("unexpected")

	t.mockSQL.ExpectBegin()
	t.mockSQL.ExpectRollback().WillReturnError(mockErr)

	err := t.db.Tx(t.ctx, func(ctx context.Context, err chan error) {
		err <- mockErr
		return
	})

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Error(err)
}

func (t *dbTestSuite) TestTx_Rollback_NoError() {
	mockErr := errors.New("unexpected")

	t.mockSQL.ExpectBegin()
	t.mockSQL.ExpectRollback()

	err := t.db.Tx(t.ctx, func(ctx context.Context, err chan error) {
		err <- mockErr
		return
	})

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Error(err)
}

func (t *dbTestSuite) TestNestedTx_Rollback_NoError() {
	mockErr := errors.New("unexpected")

	t.mockSQL.ExpectBegin()
	t.mockSQL.ExpectRollback()

	err := t.db.Tx(t.ctx, func(ctx context.Context, e chan error) {
		if err := t.db.Tx(ctx, func(ctx context.Context, e chan error) {
			e <- nil
			return
		}); err != nil {
			e <- err
			return
		}

		if err := t.db.Tx(ctx, func(ctx context.Context, e chan error) {
			e <- mockErr
			return
		}); err != nil {
			e <- err
			return
		}
	})

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Error(err)
}

func (t *dbTestSuite) TestExecContext_TxNotOk() {
	_, err := t.db.ExecContext(t.ctx, "", nil)

	t.Error(err)
}

func (t *dbTestSuite) TestExecContext_TxOk() {
	t.mockSQL.ExpectBegin()
	t.mockSQL.ExpectRollback()

	err := t.db.Tx(t.ctx, func(ctx context.Context, e chan error) {
		_, err := t.db.ExecContext(ctx, "", "")
		e <- err
		return
	})

	t.NoError(t.mockSQL.ExpectationsWereMet())
	t.Error(err)
}

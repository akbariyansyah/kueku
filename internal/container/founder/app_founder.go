package founder

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"kueku/config"
	"kueku/internal/container"
)

type app struct {
	conf *config.Config
	db   *sqlx.DB

	path string
}

// Config .
func (a *app) Config() *config.Config {
	if a.conf == nil {
		cfg, err := config.Read(a.path)
		if err != nil {
			panic(err)
		}
		a.conf = cfg
	}
	return a.conf
}

// DB .
func (a *app) DB() *sqlx.DB {
	if a.db == nil {
		db := a.buildSqlx(&a.conf.Database)

		a.db = db
	}
	return a.db
}

// NewApp .
func NewApp(path string) container.App {
	return &app{path: path}
}

func (a *app) buildSqlx(conf *config.Database) *sqlx.DB {
	db := sqlx.MustConnect(conf.Driver, conf.DSN)
	db.SetMaxIdleConns(conf.MaxIdleConn)
	db.SetMaxOpenConns(conf.MaxOpenConn)
	db.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetime) * time.Hour)
	return db
}

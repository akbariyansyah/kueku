package container

import (
	"kueku/config"
	"github.com/jmoiron/sqlx"
)

// App .
type App interface {
	Config() *config.Config
	DB() *sqlx.DB
}

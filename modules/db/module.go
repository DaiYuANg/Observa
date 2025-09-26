package db

import (
	"database/sql"

	"go.uber.org/fx"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var Module = fx.Module("db", fx.Provide(newConn))

func newConn() (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:")
	db := bun.NewDB(sqldb, sqlitedialect.New())
	return db, err
}

package seed

import (
	"database/sql"
	_ "embed"
)

//go:embed create_seed.sql
var createSeedSql string

type DB interface {
	Exec(string, ...interface{}) (sql.Result, error)
}

func MigrateTables(db DB) error {
	_, err := db.Exec(createSeedSql)
	return err
}

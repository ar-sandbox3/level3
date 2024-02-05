package level3

import (
	_ "embed"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//go:embed database/migrations/initial.sql
var initial string

// Initialize creates a new connection to the database and runs the schema.
func Initialize() (*sqlx.DB, error) {
	// Connect to the database. Please change this to your own database.
	db, err := sqlx.Connect("mysql", "root:password@tcp(localhost:3306)/level3?multiStatements=true")
	if err != nil {
		return nil, err
	}

	// We exec the initial sql to create the schema.
	db.MustExec(initial)
	return db, nil
}

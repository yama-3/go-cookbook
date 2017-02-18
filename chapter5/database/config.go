package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

// Example hold the results of our queries
type Example struct {
	Name    string
	Created *time.Time
}

// Setup configures and returns our database
// connection poold
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@/gocookbook?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

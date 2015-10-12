package testdata

import (
	"database/sql"
)

func CreateTestDate(db *sql.DB) {
	insertProducts(db)
}

package pkg

import (
	"database/sql"
	"fmt"
)

func Truncate(db *sql.DB, table string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to create a transaction: %s", err.Error())
	}
	if _, err = tx.Exec(`truncate table ` + table); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to truncate: %s", err.Error())
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit: %s", err.Error())
	}
	return nil
}

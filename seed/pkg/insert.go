package pkg

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"syreclabs.com/go/faker"
)

func Insert(db *sql.DB, table string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to create a transaction: %s", err.Error())
	}
	switch table {
	case "users":
		err = user(tx)
	default:
		err = errors.New("table does not exist")
	}
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert into %s: %s", table, err.Error())
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit: %s", err.Error())
	}
	return nil
}

func user(tx *sql.Tx) error {
	for i := 0; i <= 21; i++ {
		email := faker.Internet().Email()
		password, err := bcrypt.GenerateFromPassword([]byte("password"), 14)
		if err != nil {
			return fmt.Errorf("failed to generate password hash: %s", err.Error())
		}
		_, err = tx.Exec(
			`insert into users (email, password)
			values ($1, $2)`,
			email,
			string(password),
		)
		if err != nil {
			return fmt.Errorf("failed to insert user: %s", err.Error())
		}
	}
	return nil
}

package pkg

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DB() *sql.DB {
	url := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())

	} else if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %s", err.Error())
	}
	return db
}

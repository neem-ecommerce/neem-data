package main

import (
	"flag"
	"log"
	"neem-seed/pkg"
)

func main() {
	db := pkg.DB()
	var operation string
	var table string
	flag.StringVar(
		&operation,
		"op",
		"insert",
		"Insert records or truncate database tables. Allowed options: insert, truncate",
	)
	flag.StringVar(
		&table,
		"t",
		"",
		"Table name to insert or truncate records.",
	)
	flag.Parse()
	print(operation, " and ", table)
	defer db.Close()
	switch operation {
	case "insert":
		log.Fatal(pkg.Insert(db, table))
	case "truncate":
		log.Fatal(pkg.Truncate(db, table))
	}
}

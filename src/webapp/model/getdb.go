package model

import "database/sql"

var db *sql.DB

func HandleDB(database *sql.DB) {
	db = database
}

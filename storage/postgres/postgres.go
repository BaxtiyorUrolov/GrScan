package postgres

import "database/sql"

type Storage struct {
	DB *sql.DB
}


package database

import "database/sql"

type DatabaseConnection struct {
	*sql.DB
}

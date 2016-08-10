package database

import "database/sql"

type DatabaseConnection struct {
	*sql.DB
}

func NewDatabase(type string) DatabaseConnection {
    if type == "production" {
        //do something
    } else {
        db, err :=
    }
}

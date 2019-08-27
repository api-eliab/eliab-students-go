package main

import (
	"database/sql"
)

// School doc ...
type School struct{
	ID int64
	db *sql.DB
}

var catalog map[string]*sql.DB

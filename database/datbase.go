package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func Init(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to connect db: %v", err)
	}

	stmt := `
	create table if not exists users (
    	user_id text primary key not null,
    	display_name text default 'unknown'
	);`
	_, err = db.Exec(stmt)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	stmt = `
	create table if not exists messages(
    	message_id integer primary key autoincrement,
    	'from'     text not null,
		'to'       text not null,
    	'text' text not null default '',
    	metadata text not null default '{}',
    	kind integer not null  default 0,
    	created_at integer not null
	);`

	_, err = db.Exec(stmt)
	if err != nil {
		return fmt.Errorf("failed to create messages table: %v", err)
	}

	database = db

	return nil
}

func Close() {
	_ = database.Close()
}

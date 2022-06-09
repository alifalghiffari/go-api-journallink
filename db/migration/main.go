package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
    id integer not null primary key AUTOINCREMENT,
    username varchar(255) not null,
    password varchar(255) not null,
    role varchar(255) not null,
    created_at datetime not null
);

INSERT INTO users(username, password, role, created_at) VALUES
    ('dito', 'dito332', 'mahasiswa', '');`)

	if err != nil {
		panic(err)
	}
}

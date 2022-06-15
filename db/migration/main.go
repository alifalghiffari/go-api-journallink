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
	email varchar(255) not null,
    password varchar(255) not null,
    role "enum('mentor','mahasiswa')" not null DEFAULT 'mahasiswa',
    created_at datetime not null,
	updated_at datetime not null,
	unique(email)
);

CREATE TABLE IF NOT EXISTS journals (
    id integer not null primary key AUTOINCREMENT,
    user_id integer default null,
	isi varchar(255) not null,
	status "enum('pending','review','revisi','selesai')" not null DEFAULT 'pending',
	date_submit datetime not null,
	created_at datetime not null,
	updated_at datetime not null,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO users (
	username,
	email,
	password,
	role,
	created_at,
	updated_at
)
VALUES (
	'yap',
	'yap@gmail.com',
	'lntanasm',
	'mentor',
	'2022-06-13 22:04:59',
	'2022-06-13 22:04:59'
);

INSERT INTO journals(
	user_id,
	isi,
	status,
	date_submit,
	created_at,
	updated_at
)
VALUES (
	'4',
	'jurnalny siapa',
	'pending',
	'2022-06-13 23:35:59',
	'2022-06-13 23:35:59',
	'2022-06-13 23:35:59'
);`)

	if err != nil {
		panic(err)
	}
}

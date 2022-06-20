package repository

import (
	"database/sql"
	"errors"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {
	var sqlStmt string
	var user User

	sqlStmt = `SELECT id, username, email, password, role, created_at, updated_at FROM users WHERE id = ?;`

	row := u.db.QueryRow(sqlStmt, id)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.Created_at,
		&user.Updated_at,
	)

	return user, err
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	var sqlStmt string
	var users []User

	sqlStmt = `SELECT id, username, email, password, role, created_at, updated_at FROM users`

	rows, err := u.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Role,
			&user.Password,
			&user.Created_at,
			&user.Updated_at,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) Login(email string, password string) (*string, error) {
	var sqlStmt string

	sqlStmt = `SELECT id, username, email, password, role, created_at, updated_at FROM users WHERE email = ? AND password = ?`

	row := u.db.QueryRow(sqlStmt, email, password)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.Created_at,
		&user.Updated_at,
	)

	if err != nil {
		return nil, errors.New("Invalid username or password")
	}

	return &user.Username, nil

}

func (u *UserRepository) InsertUser(username string, email string, password string, role string) error {
	var sqlStmt string

	sqlStmt = `INSERT INTO users (username, email, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?);`

	_, err := u.db.Exec(sqlStmt, username, email, password, role, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {
	var sqlStmt string
	var role string

	sqlStmt = `SELECT role FROM users WHERE username = ?;`

	row := u.db.QueryRow(sqlStmt, username)
	err := row.Scan(&role)

	return &role, err
}

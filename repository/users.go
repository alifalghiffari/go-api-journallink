package repository

import (
	"database/sql"
	"errors"
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

	sqlStmt = `SELECT id, username, password, role, created_at FROM users WHERE id = ?;`

	row := u.db.QueryRow(sqlStmt, id)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.created_at,
	)

	return user, err
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	var sqlStmt string
	var users []User

	sqlStmt = `SELECT id, username, password, role, created_at FROM users`

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
			&user.Password,
			&user.Role,
			&user.created_at,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) Login(username string, password string) (*string, error) {
	var sqlStmt string

	sqlStmt = `SELECT id, username, password, role, created_at FROM users WHERE username = ? AND password = ?;`

	row := u.db.QueryRow(sqlStmt, username, password)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.created_at,
	)

	if err != nil {
		return nil, errors.New("Login Failed")
	}

	if user.Username == username && user.Password == password {
		sqlStmtStatus := `UPDATE users SET created_at = TRUE WHERE username = ?;`
		_, err := u.db.Exec(sqlStmtStatus, username)
		if err != nil {
			return nil, err
		}
		return &user.Username, nil
	}

	return nil, errors.New("Login Failed")

}

func (u *UserRepository) InsertUser(username string, password string, role string, created_at bool) error {
	var sqlStmt string

	sqlStmt = `INSERT INTO users (username, password, role, created_at) VALUES (?, ?, ?, ?);`

	_, err := u.db.Exec(sqlStmt, username, password, role, created_at)
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

package repository

import "time"

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Email   string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
	Token    string `db:"token"`
}

type Journal struct {
	ID       int64  `db:"id"`
	UserID   int64  `db:"user_id"`
	Isi      string `db:"isi"`
	Status   string `db:"status"`
	DateSubmit time.Time `db:"date_submit"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}

type Notification struct {
	ID       int64  `db:"id"`
	UserID   int64  `db:"user_id"`
	Message string  `db:"message"`
	Status  string `db:"status"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}

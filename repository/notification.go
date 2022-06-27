package repository

import (
	"database/sql"
	"errors"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (n *NotificationRepository) GetNotificationByUserID(id int64) (Notification, error) {
	var sqlStmt string
	var notification Notification

	sqlStmt = `SELECT id, user_id, message FROM notifications WHERE user_id = ?;`

	row := n.db.QueryRow(sqlStmt, id)
	err := row.Scan(
		&notification.ID,
		&notification.UserID,
		&notification.Message,
	)

	if err != nil {
		return notification, errors.New("No Available Journal")
	}

	return notification, nil
}

func (n *NotificationRepository) SendNotification(id int64) error {
	var sqlStmt string

	sqlStmt = `INSERT INTO notifications (id, user_id, message) VALUES (?, ?, ?);`

	_, err := n.db.Exec(sqlStmt, id)

	if err != nil {
		return err
	}

	return nil
}

// func (n *NotificationRepository) CreateNotification(notification Notification) (Notification, error) {
// 	var sqlStmt string

// 	sqlStmt = `INSERT INTO notifications (user_id, message) VALUES (?, ?);`

// 	stmt, err := n.db.Prepare(sqlStmt)
// 	if err != nil {
// 		return notification, err
// 	}

// 	res, err := stmt.Exec(notification.UserID, notification.Message)
// 	if err != nil {
// 		return notification, err
// 	}

// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return notification, err
// 	}

// 	notification.ID = id

// 	return notification, nil
// }
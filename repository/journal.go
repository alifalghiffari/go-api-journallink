package repository

import (
	"database/sql"
	"errors"
	"time"
)

type JournalRepository struct {
	db *sql.DB
}

func NewJournalRepository(db *sql.DB) *JournalRepository {
	return &JournalRepository{db: db}
}

func (j *JournalRepository) FetchJournals() ([]Journal, error) {
	var sqlStatement string
	var journals []Journal

	sqlStatement = `SELECT id, user_id, isi, status, date_submit, created_at, updated_at FROM journals`

	rows, err := j.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	var journal Journal
	for rows.Next() {
		err := rows.Scan(
			&journal.ID,
			&journal.UserID,
			&journal.Isi,
			&journal.Status,
			&journal.DateSubmit,
			&journal.Created_at,
			&journal.Updated_at,
		)

		if err != nil {
			return nil, errors.New("No Available Journal")
		}
		journals = append(journals, journal)
	}

	return journals, err
}

func (j *JournalRepository) GetJournalByID(id int64) (Journal, error) {
	var sqlStatement string
	var journal Journal

	sqlStatement = `SELECT id, user_id, isi, status, date_submit, created_at, updated_at FROM journals WHERE id = ?`

	err := j.db.QueryRow(sqlStatement, id).Scan(
		&journal.ID,
		&journal.UserID,
		&journal.Isi,
		&journal.Status,
		&journal.DateSubmit,
		&journal.Created_at,
		&journal.Updated_at,
	)

	if err != nil {
		return journal, errors.New("No Available Journal")
	}

	return journal, nil
}

func (j *JournalRepository) InsertJournal(UserID int64, Isi string, Status string, DateSubmit string) error {
	var sqlStmt string

	sqlStmt = `INSERT INTO journals (user_id, isi, status, date_submit, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?);`

	_, err := j.db.Exec(sqlStmt, UserID, Isi, Status, time.Now(), time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (j *JournalRepository) UpdateJournal(Isi string, id int64, user_id int64) error {
	var sqlStmt string

	sqlStmt = `UPDATE journals SET isi = ? WHERE id = ? AND user_id =?;`

	_, err := j.db.Exec(sqlStmt, Isi, id, user_id)

	if err != nil {
		return err
	}

	return nil
}

func (j *JournalRepository) UpdateStatus(Status string, id int64) error {
	var sqlStmt string

	sqlStmt = `UPDATE journals SET status = ? WHERE id = ?;`

	_, err := j.db.Exec(sqlStmt, Status, id)

	if err != nil {
		return err
	}

	return nil
}
package main

import (
	"database/sql"

	"go-api-project/api"
	"go-api-project/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)
	journalRepo := repository.NewJournalRepository(db)
	notificationRepo := repository.NewNotificationRepository(db)

	mainAPI := api.NewAPI(*usersRepo, *journalRepo, *notificationRepo)
	mainAPI.Start()
}

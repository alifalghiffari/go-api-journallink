package repository_test

import (
	"database/sql"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	_ "github.com/mattn/go-sqlite3"

	"go-api-project/repository"
)

var _ = Describe("Repository Test", func() {
	var db *sql.DB
	var err error
	var usersRepo *repository.UserRepository

	BeforeEach(func() {
		// Setup
		db, err = sql.Open("sqlite3", "./testing.db")
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

		
		INSERT INTO users(username, email, password, role, created_at, updated_at) VALUES
			('aditira', 'abc@gmail.com', '$2a$10$oth77hKkP4VBHjOjZy1QteoT52G4POLvV.T0teKSI/7uSh55n6Kxi', 'mentor', '2022-01-05 17:10:16', '2022-01-09 17:10:16'),
			('dina', 'abcd@gmail.com', '$2a$10$2Qwv5tE672Ukpu0idh0XxuT/prIoYQa3tikLo2UXcJo8GH5Aizbti', 'mahasiswa', '2022-02-05 17:05:16', '2022-01-11 17:10:16'),
			('dito', 'abce@gmail.com', '$2a$10$DQp6JNtDDP8dQS3xSjFu5ufYckFo8hDzcYHaItRcCglH8P.glco2q', 'mahasiswa', '2022-03-05 17:05:16', '2022-01-22 17:10:16');`)

		if err != nil {
			panic(err)
		}

		usersRepo = repository.NewUserRepository(db)
	})

	AfterEach(func() {
		// Teardown
		db, err := sql.Open("sqlite3", "./testing.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		DROP TABLE IF EXISTS users;`)

		if err != nil {
			panic(err)
		}
	})

	Describe("Login", func() {
		When("Email and password are correct", func() {
			It("accepts the login", func() {
				_, err := usersRepo.Login("abc@gmail.com", "$2a$10$oth77hKkP4VBHjOjZy1QteoT52G4POLvV.T0teKSI/7uSh55n6Kxi")
				Expect(err).ToNot(HaveOccurred())
			})
		})
		When("username is correct but password is incorrect", func() {
			It("rejects the login", func() {
				_, err := usersRepo.Login("abc@gmail.com", "4567")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Invalid Email or Password"))
			})
		})
		When("both username and password is incorrect", func() {
			It("rejects the login", func() {
				_, err := usersRepo.Login("abcz@gmail.com", "23885")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

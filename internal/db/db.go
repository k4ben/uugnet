package db

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const dbPath = "sqlite.db"

type User struct {
	Username string
	Password string
}

type UserDbRow struct {
	ID int
	User
}

func InitDatabase() error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS User (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username varchar(255) UNIQUE,
			password varchar(255)
		)`,
	)
	if err != nil {
		return err
	}
	return nil
}

func AddUser(u *User) error {
	_, err := db.ExecContext(
		context.Background(),
		`INSERT INTO User (username, password) VALUES (?,?);`, u.Username, u.Password,
	)
	return err
}

func DelUser(username string) error {
	_, err := db.ExecContext(
		context.Background(),
		`DELETE FROM User WHERE username=?;`, username,
	)
	return err
}

func GetUser(username string) (UserDbRow, error) {
	var user UserDbRow

	row := db.QueryRowContext(
		context.Background(),
		`SELECT * FROM User WHERE username=?;`, username,
	)

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

func GetUsers() ([]UserDbRow, error) {
	var users []UserDbRow
	rows, err := db.QueryContext(
		context.Background(),
		`SELECT * FROM User;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user UserDbRow
		if err := rows.Scan(
			&user.ID, &user.Username, &user.Password,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

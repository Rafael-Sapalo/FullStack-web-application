package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func IsEmailAlreadyExist(db *sql.DB, email string) (bool, error) {
	var qr = "SELECT COUNT(*) FROM users WHERE email = ?"
	statement, err := db.Prepare(qr)
	if err != nil {
		return false, nil
	}
	defer statement.Close()

	var count int
	err = statement.QueryRow(email).Scan(&count)
	if err != nil {
		return false, nil
	}

	return count > 0, nil
}

func IsUsernameAlreadyExist(db *sql.DB, username string) (bool, error) {
	var qr = "SELECT COUNT(*) FROM users WHERE username = ?"
	var statement, err = db.Prepare(qr)
	if err != nil {
		return false, nil
	}
	defer statement.Close()
	var count int
	err = statement.QueryRow(username).Scan(&count)
	if err != nil {
		return false, nil
	}
	return count > 0, nil
}

func GetRoleWthId(db *sql.DB, userID int) (string, error) {
	var qr string = "SELECT roles FROM users WHERE id = ?"
	var roles string
	var trans, TrErr = db.Begin()
	if TrErr != nil {
		return "", TrErr
	}
	defer trans.Rollback()
	if err := trans.QueryRow(qr, userID).Scan(&roles); err != nil {
		return "", err
	}
	if err := trans.Commit(); err != nil {
		return "", err
	}
	return roles, nil
}

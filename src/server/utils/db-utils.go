package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func IsEmailAlreadyExist(db *sql.DB, email string) (bool, error) {
	var qr = "SELECT COUNT(*) FROM users WHERE email = ?";
	statement, err := db.Prepare(qr);
	if err != nil {
		return false, nil;
	}
	defer statement.Close();

	var count int;
	err = statement.QueryRow(email).Scan(&count);
	if err != nil {
		return false, nil;
	}
	
	return count > 0, nil;
}
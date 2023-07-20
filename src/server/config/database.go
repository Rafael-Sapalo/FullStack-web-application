package config

import (
	"fmt"
	"time"

	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/TwiN/go-color"
	"github.com/briandowns/spinner"
)

func ConnectDB() (*sql.DB, error) {
	s := spinner.New(spinner.CharSets[36], 415*time.Millisecond);
	s.Prefix = "Connecting to the database ";
	s.Start();
	db, err := sql.Open("mysql", "rafael:2da89c775@tcp(127.0.0.1:3306)/FullStack");

	if err != nil {
		return nil, err;
	}
	err = db.Ping();
	time.Sleep(5 * time.Second);
	if err != nil {
		s.Stop();
		fmt.Println(color.With(color.Red, "Failed to connect to the database"));
		db.Close();
		return nil, err;
	}
	s.Stop();
	fmt.Println(color.With(color.Green, "Successfully connected to the database"));

	return db, nil;
}

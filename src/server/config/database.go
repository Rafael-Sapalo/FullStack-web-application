package config

import (
	"fmt"
	"os"
	"time"

	"database/sql"

	"github.com/TwiN/go-color"
	"github.com/briandowns/spinner"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func AttachDB(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}

func GetDBsource() any {

	var err = godotenv.Load()
	if err != nil {
		fmt.Println(color.With(color.Red, "No env specified or found"))
	}
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbUser = os.Getenv("DB_USER")
	var dbPass = os.Getenv("DB_PASS")
	var dbName = os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" {
		fmt.Println(color.With(color.Red, "Missing database environment variables"))
		return nil
	}
	var dbSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	return dbSource
}

func ConnectDB() (*sql.DB, error) {

	var dbSource = GetDBsource()
	if dbSource == nil {
		return nil, nil
	}
	var sp = spinner.New(spinner.CharSets[36], 415*time.Millisecond)
	sp.Prefix = "Connecting to the database "
	sp.Start()
	db, err := sql.Open("mysql", dbSource.(string))

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	time.Sleep(5 * time.Second)
	if err != nil {
		sp.Stop()
		fmt.Println(color.With(color.Red, "Failed to connect to the database"))
		db.Close()
		return nil, err
	}
	sp.Stop()
	fmt.Println(color.With(color.Green, "Successfully connected to the database"))

	return db, nil
}

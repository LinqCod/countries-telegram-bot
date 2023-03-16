package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func InitDB() (*sql.DB, error) {
	host := viper.GetString("HOST")
	port := viper.GetString("PORT")
	username := viper.GetString("USERNAME")
	password := viper.GetString("PASSWORD")
	dbname := viper.GetString("DBNAME")

	pgInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username, password, host, port, dbname)

	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		return nil, fmt.Errorf("validation of db parameters failed due to error: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to open db connection due to err: %v", err)
	}

	log.Println("postgres db connected successfully!")
	return db, nil
}

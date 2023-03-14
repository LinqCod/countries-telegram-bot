package database

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func InitDB() (*sql.DB, error) {
	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.port")
	username := viper.GetString("postgres.username")
	password := viper.GetString("postgres.password")
	dbname := viper.GetString("postgres.dbname")

	pgInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=false",
		host, port, username, password, dbname)

	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		return nil, fmt.Errorf("validation of db parameters failed due to apierrors: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to open db connection due to err: %v", err)
	}

	log.Println("postgres db connected successfully!")
	return db, nil
}

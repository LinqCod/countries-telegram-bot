package repository

import (
	"database/sql"
	"github.com/linqcod/countries-telegram-bot/internal/model"
	"log"
)

type CountriesRepository struct {
	db *sql.DB
}

func NewCountriesRepository(db *sql.DB) *CountriesRepository {
	return &CountriesRepository{
		db: db,
	}
}

func (r *CountriesRepository) SaveCountry(country *model.Country) error {
	log.Println("Saving country to db")
	_, err := r.db.Exec("INSERT INTO country (information) VALUES($1)", country)
	if err != nil {
		return err
	}
	log.Println("country saved successfully!")

	return nil
}

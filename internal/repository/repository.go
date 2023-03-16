package repository

import (
	"database/sql"
	"github.com/linqcod/countries-telegram-bot/internal/model"
)

type CountriesRepository struct {
	db *sql.DB
}

func NewCountriesRepository(db *sql.DB) *CountriesRepository {
	return &CountriesRepository{
		db: db,
	}
}

func (r *CountriesRepository) SaveCountry(country *model.Country) {
	//TODO: save country to db
}

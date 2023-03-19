package repository

import (
	"database/sql"
	"github.com/linqcod/countries-telegram-bot/internal/countries/model"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveCountry(country *model.Country) error {
	log.Println("Saving country to db")
	_, err := r.db.Exec(`INSERT INTO country 
    (
    	name, 
        official_name, 
        is_independent,
        status,
     	is_un_member,
     	region,	
     	sub_region,
     	area,
     	population
    ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		country.Name.Common,
		country.Name.Official,
		country.Independent,
		country.Status,
		country.UnMember,
		country.Region,
		country.SubRegion,
		country.Area,
		country.Population,
	)
	if err != nil {
		return err
	}
	log.Println("country saved successfully!")

	return nil
}

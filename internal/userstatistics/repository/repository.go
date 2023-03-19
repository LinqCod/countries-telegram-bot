package repository

import (
	"database/sql"
	countryModel "github.com/linqcod/countries-telegram-bot/internal/countries/model"
	statsModel "github.com/linqcod/countries-telegram-bot/internal/userstatistics/model"
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

func (r *Repository) GetStatistics() (*statsModel.Statistics, error) {
	log.Println("getting all statistics from db")

	countryNameByArea, err := r.getGreatestCountryByArea()
	if err != nil {
		return nil, err
	}
	countryNameByPopulation, err := r.getGreatestCountryByPopulation()
	if err != nil {
		return nil, err
	}
	region, err := r.getMostFrequentCountryRegion()
	if err != nil {
		return nil, err
	}

	log.Println("all statistics successfully got!")

	return &statsModel.Statistics{
		GreatestCountryByArea:       countryNameByArea,
		GreatestCountryByPopulation: countryNameByPopulation,
		MostFrequentCountryRegion:   region,
	}, nil
}

func (r *Repository) getGreatestCountryByArea() (countryModel.Name, error) {
	log.Println("getting greatest country by area from db")

	row := r.db.QueryRow(`SELECT name, official_name FROM country ORDER BY area DESC LIMIT 1`)

	countryName := countryModel.Name{}
	if err := row.Scan(&countryName.Common, &countryName.Official); err != nil {
		return countryName, err
	}

	log.Println("country successfully got!")

	return countryName, nil
}

func (r *Repository) getGreatestCountryByPopulation() (countryModel.Name, error) {
	log.Println("getting greatest country by population from db")

	row := r.db.QueryRow(`SELECT name, official_name FROM country ORDER BY population DESC LIMIT 1`)

	countryName := countryModel.Name{}
	if err := row.Scan(&countryName.Common, &countryName.Official); err != nil {
		return countryName, err
	}

	log.Println("country successfully got!")

	return countryName, nil
}

func (r *Repository) getMostFrequentCountryRegion() (string, error) {
	log.Println("getting most frequent country region from db")

	row := r.db.QueryRow(`SELECT region FROM country GROUP BY region ORDER BY count(*) DESC LIMIT 1`)

	var region string
	if err := row.Scan(&region); err != nil {
		return "", err
	}

	log.Println("region successfully got!")

	return region, nil
}

package model

import "github.com/linqcod/countries-telegram-bot/internal/countries/model"

type Statistics struct {
	GreatestCountryByArea       model.Name
	GreatestCountryByPopulation model.Name
	MostFrequentCountryRegion   string
}

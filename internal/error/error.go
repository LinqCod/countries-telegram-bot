package error

import "errors"

var CountryNotFound = errors.New("country not found")

type RestCountriesApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

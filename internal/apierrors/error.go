package apierrors

import "errors"

var ErrorCountryNotFound = errors.New("country not found")

type RestCountriesApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

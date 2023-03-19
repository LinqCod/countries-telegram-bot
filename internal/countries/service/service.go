package service

import (
	"encoding/json"
	"github.com/linqcod/countries-telegram-bot/internal/countries/model"
	"github.com/linqcod/countries-telegram-bot/internal/countries/repository"
	apierror "github.com/linqcod/countries-telegram-bot/internal/errors"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type Service struct {
	baseUrl    string
	repository *repository.Repository
}

func NewRestCountriesService(repository *repository.Repository) *Service {
	baseUrl := viper.GetString("BASE_URL")

	return &Service{
		baseUrl:    baseUrl,
		repository: repository,
	}
}

func (a *Service) GetCountryByName(name string) (*model.Country, error) {
	response, err := http.Get(a.baseUrl + "name/" + name)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := make([]*model.Country, 1)
	if err = json.Unmarshal(content, &res); err != nil {
		apiError := apierror.RestCountriesApiError{}
		err = json.Unmarshal(content, &apiError)
		if err == nil && apiError.Message == "Not Found" {
			return nil, apierror.CountryNotFound
		}

		return nil, err
	}

	if err = a.repository.SaveCountry(res[0]); err != nil {
		return nil, err
	}

	return res[0], nil
}

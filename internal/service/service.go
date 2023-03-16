package service

import (
	"encoding/json"
	apierror "github.com/linqcod/countries-telegram-bot/internal/error"
	"github.com/linqcod/countries-telegram-bot/internal/model"
	"github.com/linqcod/countries-telegram-bot/internal/repository"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type RestCountriesService struct {
	baseUrl    string
	repository *repository.CountriesRepository
}

func NewRestCountriesService(repository *repository.CountriesRepository) *RestCountriesService {
	baseUrl := viper.GetString("BASE_URL")

	return &RestCountriesService{
		baseUrl:    baseUrl,
		repository: repository,
	}
}

func (a *RestCountriesService) GetCountryByName(name string) (*model.Country, error) {
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

	//TODO: repository.SaveCountry

	return res[0], nil
}

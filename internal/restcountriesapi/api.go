package restcountriesapi

import (
	"encoding/json"
	"github.com/linqcod/countries-telegram-bot/internal/apierrors"
	"github.com/linqcod/countries-telegram-bot/internal/model"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type CountriesApi struct {
	baseUrl string
}

func NewCountriesApi() *CountriesApi {
	baseUrl := viper.GetString("rest-countries-api.baseUrl")

	return &CountriesApi{
		baseUrl: baseUrl,
	}
}

func (a *CountriesApi) GetCountryByName(name string) (*model.Country, error) {
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
		apiError := apierrors.RestCountriesApiError{}
		err = json.Unmarshal(content, &apiError)
		if err == nil && apiError.Message == "Not Found" {
			return nil, apierrors.ErrorCountryNotFound
		}

		return nil, err
	}

	return res[0], nil
}

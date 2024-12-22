package services

import (
	"github.com/rayaanshahid/restcountries/internal/api"
	"github.com/rayaanshahid/restcountries/internal/models"
)

type CurrencyService struct {
	client *api.Client
}

func NewCurrencyService(client *api.Client) *CurrencyService {
	return &CurrencyService{client: client}
}

func (s *CurrencyService) GetAllCurrencies() ([]models.CurrencyDto, error) {
	countries, err := s.client.FetchCountries()
	if err != nil {
		return nil, err
	}

	currencyMap := map[string][]string{}
	for _, country := range countries {
		for code := range country.Currencies {
			currencyMap[code] = append(currencyMap[code], country.CountryName.Common)
		}
	}

	currencies := []models.CurrencyDto{}
	for code, countryList := range currencyMap {
		currencies = append(currencies, models.CurrencyDto{
			Name:      code,
			Countries: countryList,
		})
	}

	return currencies, nil
}

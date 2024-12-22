package services

import (
	"github.com/rayaanshahid/restcountries/internal/api"
	"github.com/rayaanshahid/restcountries/internal/models"
)

type CountryService struct {
	client *api.Client
}

func NewCountryService(client *api.Client) *CountryService {
	return &CountryService{client: client}
}

func (s *CountryService) GetEuropeanCountries() ([]models.CountryDto, error) {
	countries, err := s.client.FetchEuropeanCountries()
	if err != nil {
		return nil, err
	}

	europeanCountries := []models.CountryDto{}
	for _, country := range countries {
		if country.Region == "Europe" {
			for code, currency := range country.Currencies {
				europeanCountries = append(europeanCountries, models.CountryDto{
					Name:     country.CountryName.Common,
					Currency: code + " (" + currency.Name + ")",
				})
			}
		}
	}

	return europeanCountries, nil
}

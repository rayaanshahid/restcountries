package tests

import (
	"testing"

	"github.com/rayaanshahid/restcountries/internal/api"
	"github.com/rayaanshahid/restcountries/internal/services"
)

func TestGetEuropeanCountries(t *testing.T) {
	client := api.NewClient("https://restcountries.com/v3.1")
	countryService := services.NewCountryService(client)

	result, err := countryService.GetEuropeanCountries()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result) == 0 {
		t.Error("Expected countries but got none")
	}
}

func TestGetAllCurrencies(t *testing.T) {
	client := api.NewClient("https://restcountries.com/v3.1")
	currencyService := services.NewCurrencyService(client)

	result, err := currencyService.GetAllCurrencies()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(result) == 0 {
		t.Error("Expected currencies but got none")
	}
}

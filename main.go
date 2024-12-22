package main

import (
	"fmt"
	"log"

	"github.com/rayaanshahid/restcountries/internal/api"
	"github.com/rayaanshahid/restcountries/internal/services"
	"github.com/rayaanshahid/restcountries/internal/utils"
)

func main() {
	client := api.NewClient("https://restcountries.com/v3.1")

	countryService := services.NewCountryService(client)
	currencyService := services.NewCurrencyService(client)

	var order = "asc"
	// Fetch and process data
	europeanCountries, err := countryService.GetEuropeanCountries()
	if err != nil {
		log.Fatalf("Error fetching European countries: %v", err)
	}

	//sort countries
	utils.SortCountries(europeanCountries, order)

	currencies, err := currencyService.GetAllCurrencies()
	if err != nil {
		log.Fatalf("Error fetching currencies: %v", err)
	}

	// Print results
	fmt.Println("European Countries:")
	for _, country := range europeanCountries {
		fmt.Printf("- %s (%s)\n", country.Name, country.Currency)
	}

	fmt.Println("\nWorld Currencies:")
	for _, currency := range currencies {
		fmt.Printf("- %s: %v\n", currency.Name, currency.Countries)
	}
}

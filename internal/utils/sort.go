package utils

import (
	"sort"

	"github.com/rayaanshahid/restcountries/internal/models"
)

// SortCountries sorts a list of countries based on their Name field
// in ascending or descending order.
func SortCountries(countries []models.CountryDto, order string) {
	if order == "asc" {
		sort.Slice(countries, func(i, j int) bool {
			return countries[i].Name < countries[j].Name
		})
	} else {
		sort.Slice(countries, func(i, j int) bool {
			return countries[i].Name > countries[j].Name
		})
	}
}

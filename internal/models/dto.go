package models

type CountryDto struct {
	Name     string `json:"name"`
	Currency string `json:"currency"`
}

type CurrencyDto struct {
	Name      string   `json:"name"`
	Countries []string `json:"countries"`
}

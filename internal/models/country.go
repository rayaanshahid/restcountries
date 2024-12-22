package models

type Name struct {
	Common     string `json:"common"`
	Official   string `json:"official"`
	NativeName map[string]struct {
		Official string `json:"official"`
		Common   string `json:"common"`
	} `json:"nativeName"`
}

type Country struct {
	CountryName Name                `json:"name"`
	Currencies  map[string]Currency `json:"currencies"`
	Region      string              `json:"region"`
}

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

package model

type SolWeather struct {
}

type WeatherWrapper struct {
	SolKeys        []string     `json:"sol_keys"`
	ValidityChecks []SolWeather `json:"validity_checks"`
}

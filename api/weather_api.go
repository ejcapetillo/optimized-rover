package api

import "errors"

type WeatherAPI interface {
	GetMarsWeather(nasaUrl string) error
}

type weatherAPI struct {
}

func NewWeatherAPI() WeatherAPI {
	return &weatherAPI{}
}

func (api *weatherAPI) GetMarsWeather(nasaUrl string) error {
	if nasaUrl == "" {
		return errors.New("missing image GET URL")
	}

	return nil
}

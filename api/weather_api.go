package api

import (
	"errors"
	"fmt"
	"net/http"
)

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
		return errors.New("missing weather GET URL")
	}

	response, err := http.Get(nasaUrl)
	if err != nil {
		return fmt.Errorf("error on weather GET request: %w", err)
	}

	if response.Body == nil {
		return fmt.Errorf("no body returned from weather GET request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unsuccessful request to weather GET API: %d", response.StatusCode)
	}

	return nil
}

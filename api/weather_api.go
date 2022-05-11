package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ejcapetillo/optimized-rover/model"
	"io"
	"net/http"
)

type WeatherAPI interface {
	GetMarsWeather(nasaUrl string) (model.WeatherWrapper, error)
}

type weatherAPI struct {
}

func NewWeatherAPI() WeatherAPI {
	return &weatherAPI{}
}

func (api *weatherAPI) GetMarsWeather(nasaUrl string) (model.WeatherWrapper, error) {
	weatherWrapper := model.WeatherWrapper{}

	if nasaUrl == "" {
		return weatherWrapper, errors.New("missing weather GET URL")
	}

	response, err := http.Get(nasaUrl)
	if err != nil {
		return weatherWrapper, fmt.Errorf("error on weather GET request: %w", err)
	}

	if response.Body == nil {
		return weatherWrapper, fmt.Errorf("no body returned from weather GET request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return weatherWrapper, fmt.Errorf("unsuccessful request to weather GET API: %d", response.StatusCode)
	}

	responseBody, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(responseBody, &weatherWrapper)
	if err != nil {
		return weatherWrapper, fmt.Errorf("error unmarshalling weather GET response: %w", err)
	}

	return weatherWrapper, nil
}

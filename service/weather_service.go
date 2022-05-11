package service

import (
	"fmt"
	"github.com/ejcapetillo/optimized-rover/api"
	"github.com/ejcapetillo/optimized-rover/model"
	"net/url"
)

type WeatherService interface {
	GetMarsWeather() (model.WeatherWrapper, error)
}

type weatherService struct {
	weatherAPI api.WeatherAPI
}

func NewWeatherService(weatherAPI api.WeatherAPI) WeatherService {
	return &weatherService{
		weatherAPI: weatherAPI,
	}
}

var nasaWeatherAPIRoot = "https://api.nasa.gov/insight_weather/?"

func (service *weatherService) GetMarsWeather() (model.WeatherWrapper, error) {
	params := url.Values{}
	params.Add("api_key", model.NasaAPIKey)
	params.Add("feedtype", "json")
	params.Add("ver", "1.0")

	nasaURL := fmt.Sprintf("%s%s", nasaWeatherAPIRoot, params.Encode())
	return service.weatherAPI.GetMarsWeather(nasaURL)
}

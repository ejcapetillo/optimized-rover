package service

import "github.com/ejcapetillo/optimized-rover/api"

type WeatherService interface {
	GetMarsWeather() error
}

type weatherService struct {
	weatherAPI api.WeatherAPI
}

func NewWeatherService(weatherAPI api.WeatherAPI) WeatherService {
	return &weatherService{
		weatherAPI: weatherAPI,
	}
}

func (service *weatherService) GetMarsWeather() error {
	return nil
}

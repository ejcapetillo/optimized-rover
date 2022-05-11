package service

import (
	"fmt"
	"github.com/ejcapetillo/optimized-rover/api"
	"github.com/ejcapetillo/optimized-rover/model"
	"net/url"
	"time"
)

type ImageService interface {
	GetImages() (map[string]int, error)
}

type imageService struct {
	imageAPI api.ImageAPI
}

func NewImageService(imageAPI api.ImageAPI) ImageService {
	return &imageService{
		imageAPI: imageAPI,
	}
}

// TODO Split into separate config or secret
var nasaAPIKey = "IWo1dIukz6jnuWiEWbl86cFnxjyOmallYO2raOnB"
var nasaAPIRoot = "https://api.nasa.gov/mars-photos/api/v1/rovers"
var numDaysToCompare = 7

func (service *imageService) GetImages() (map[string]int, error) {
	var err error
	rovers := model.GetRovers()
	imageMap := initializeImageMap()

	for dateStr := range imageMap {
		dateTotal := 0
		for _, rover := range rovers {
			params := url.Values{}
			params.Add("earth_date", dateStr)
			params.Add("api_key", nasaAPIKey)

			nasaURL := fmt.Sprintf("%s/%s/photos?%s", nasaAPIRoot, rover, params.Encode())
			images, err := service.imageAPI.GetImages(nasaURL)
			if err != nil {
				return imageMap, err
			}
			dateTotal += len(images.Photos)
		}
		imageMap[dateStr] = dateTotal
	}

	return imageMap, err
}

func initializeImageMap() map[string]int {
	imageMap := make(map[string]int)
	now := time.Now()
	day := time.Now().AddDate(0, 0, -1*numDaysToCompare)

	for day.Before(now) {
		imageMap[day.Format("2006-01-02")] = 0
		day = day.AddDate(0, 0, 1)
	}

	return imageMap
}

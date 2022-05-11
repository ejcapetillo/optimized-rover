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

var nasaRoverAPIRoot = "https://api.nasa.gov/mars-photos/api/v1/rovers"
var numDaysToCompare = 7

func (service *imageService) GetImages() (map[string]int, error) {
	var err error
	imageMap := initializeImageMap()

	for dateStr := range imageMap {
		params := url.Values{}
		params.Add("earth_date", dateStr)
		params.Add("api_key", model.NasaAPIKey)

		nasaURL := fmt.Sprintf("%s/%s/photos?%s", nasaRoverAPIRoot, model.Curiosity, params.Encode())
		images, err := service.imageAPI.GetImages(nasaURL)
		if err != nil {
			return imageMap, err
		}
		imageMap[dateStr] = len(images.Photos)
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

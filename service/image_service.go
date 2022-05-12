package service

import (
	"fmt"
	"github.com/ejcapetillo/optimized-rover/api"
	"github.com/ejcapetillo/optimized-rover/model"
	"net/url"
	"sync"
	"time"
)

type ImageService interface {
	GetImages() ([]model.DailyPhoto, error)
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

func (service *imageService) GetImages() ([]model.DailyPhoto, error) {
	imageMap := initializeImageMap()
	var dayList []model.DailyPhoto

	var wg sync.WaitGroup
	dayChan := make(chan model.DailyPhoto)

	for key := range imageMap {
		link := getNasaUrl(key)
		wg.Add(1)
		go func(link string, key string) {
			defer wg.Done()
			images, _ := service.imageAPI.GetImages(link)
			dayChan <- model.DailyPhoto{EarthDate: key, Count: len(images.Photos)}
		}(link, key)
	}

	go func() {
		wg.Wait()
		close(dayChan)
	}()

	for day := range dayChan {
		dayList = append(dayList, day)
	}

	return dayList, nil
}

func getNasaUrl(dateStr string) string {
	params := url.Values{}
	params.Add("earth_date", dateStr)
	params.Add("api_key", model.NasaAPIKey)

	nasaURL := fmt.Sprintf("%s/%s/photos?%s", nasaRoverAPIRoot, model.Curiosity, params.Encode())

	return nasaURL
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

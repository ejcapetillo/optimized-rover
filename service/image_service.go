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
	GetImagesSingleThread() ([]model.DailyPhoto, error)
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
var numDaysToCompare = 25

func (service *imageService) GetImages() ([]model.DailyPhoto, error) {
	dateList := initializeDateList()
	var dayList []model.DailyPhoto

	var wg sync.WaitGroup
	dayChan := make(chan model.DailyPhoto)

	for _, date := range dateList {
		link := getNasaUrl(date)
		wg.Add(1)
		go func(link string, date string) {
			defer wg.Done()
			images, _ := service.imageAPI.GetImages(link)
			dayChan <- model.DailyPhoto{EarthDate: date, Count: len(images.Photos)}
		}(link, date)
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

func (service *imageService) GetImagesSingleThread() ([]model.DailyPhoto, error) {
	dateList := initializeDateList()
	var dayList []model.DailyPhoto

	for _, dateStr := range dateList {
		params := url.Values{}
		params.Add("earth_date", dateStr)
		params.Add("api_key", model.NasaAPIKey)

		nasaURL := fmt.Sprintf("%s/%s/photos?%s", nasaRoverAPIRoot, model.Curiosity, params.Encode())
		images, err := service.imageAPI.GetImages(nasaURL)
		if err != nil {
			return nil, err
		}
		dayCount := model.DailyPhoto{EarthDate: dateStr, Count: len(images.Photos)}
		dayList = append(dayList, dayCount)
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

func initializeDateList() []string {
	dateList := make([]string, 0, numDaysToCompare)
	now := time.Now()
	day := time.Now().AddDate(0, 0, -1*numDaysToCompare)

	for day.Before(now) {
		dateList = append(dateList, day.Format("2006-01-02"))
		day = day.AddDate(0, 0, 1)
	}

	return dateList
}

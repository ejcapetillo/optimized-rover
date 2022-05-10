package service

import (
	"fmt"
	"github.com/ejcapetillo/optimized-rover/api"
	"github.com/ejcapetillo/optimized-rover/model"
	"net/url"
	"time"
)

type ImageService interface {
	GetImages() error
}

type imageService struct {
	imageAPI api.ImageAPI
}

func NewImageService(imageAPI api.ImageAPI) ImageService {
	return &imageService{
		imageAPI: imageAPI,
	}
}

var nasaAPIRoot = "https://api.nasa.gov/mars-photos/api/v1/rovers"

func (service *imageService) GetImages() error {
	var err error
	rovers := model.GetRovers()

	for _, rover := range rovers {
		params := url.Values{}
		params.Add("earth_date", time.Now().Format("2006-01-02"))
		params.Add("api_key", "DEMO_KEY")

		nasaURL := fmt.Sprintf("%s/%s/photos?%s", nasaAPIRoot, rover, params.Encode())
		_, err = service.imageAPI.GetImages(nasaURL)
	}

	return err
}

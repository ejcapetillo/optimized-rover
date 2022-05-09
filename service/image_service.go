package service

import "github.com/ejcapetillo/optimized-rover/api"

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

func (service *imageService) GetImages() error {
	return service.imageAPI.GetImages()
}

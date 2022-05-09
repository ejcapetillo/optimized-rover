package main

import (
	"github.com/ejcapetillo/optimized-rover/api"
	"github.com/ejcapetillo/optimized-rover/service"
)

func main() {
	imageService := service.NewImageService(api.NewImageAPI())
	_ = imageService.GetImages()
}

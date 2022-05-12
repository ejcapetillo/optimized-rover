package main

import (
	"fmt"
	"github.com/ejcapetillo/optimized-rover/api"
	"github.com/ejcapetillo/optimized-rover/service"
	"time"
)

func main() {
	_ = service.NewWeatherService(api.NewWeatherAPI())

	imageStart := time.Now()
	imageService := service.NewImageService(api.NewImageAPI())
	imageList, err := imageService.GetImages()
	if err != nil {
		fmt.Printf("error found: %s", err)
		return
	}
	imageElapsed := time.Since(imageStart)
	fmt.Printf("Image retrieval took %s\n", imageElapsed)

	for _, day := range imageList {
		fmt.Printf("%d images taken by Curiosity on %s\n", day.Count, day.EarthDate)
	}
}

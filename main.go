package main

import (
	"fmt"
	"github.com/ejcapetillo/optimized-rover/api"
	"github.com/ejcapetillo/optimized-rover/service"
	"time"
)

func main() {
	_ = service.NewWeatherService(api.NewWeatherAPI())
	imageService := service.NewImageService(api.NewImageAPI())

	singleStart := time.Now()
	singleImageList, err := imageService.GetImagesSingleThread()
	if err != nil {
		fmt.Printf("error found: %s", err)
		return
	}
	singleElapsed := time.Since(singleStart)
	fmt.Printf("GetImagesSingleThread: Image retrieval took %s\n", singleElapsed)
	for _, day := range singleImageList {
		fmt.Printf("%d images taken by Curiosity on %s\n", day.Count, day.EarthDate)
	}

	imageStart := time.Now()
	imageList, err := imageService.GetImages()
	if err != nil {
		fmt.Printf("error found: %s", err)
		return
	}
	imageElapsed := time.Since(imageStart)
	fmt.Printf("\nGetImages: Image retrieval took %s\n", imageElapsed)

	for _, day := range imageList {
		fmt.Printf("%d images taken by Curiosity on %s\n", day.Count, day.EarthDate)
	}
}

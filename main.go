package main

import (
	"fmt"
	"github.com/ejcapetillo/optimized-rover/api"
	"github.com/ejcapetillo/optimized-rover/service"
)

func main() {
	imageService := service.NewImageService(api.NewImageAPI())
	imageMap, err := imageService.GetImages()
	if err != nil {
		fmt.Printf("error found: %s", err)
		return
	}

	for key, value := range imageMap {
		fmt.Printf("%s had %d images taken by all rovers\n", key, value)
	}
}

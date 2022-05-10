package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ejcapetillo/optimized-rover/model"
	"io"
	"net/http"
)

type ImageAPI interface {
	GetImages(nasaUrl string) (model.PhotoWrapper, error)
}

type imageAPI struct {
}

func NewImageAPI() ImageAPI {
	return &imageAPI{}
}

func (api *imageAPI) GetImages(nasaUrl string) (model.PhotoWrapper, error) {
	photos := model.PhotoWrapper{}

	if nasaUrl == "" {
		return photos, errors.New("missing image GET URL")
	}

	response, err := http.Get(nasaUrl)
	if err != nil {
		return photos, fmt.Errorf("error on image GET request: %w", err)
	}

	if response.Body == nil {
		return photos, fmt.Errorf("no body returned from image GET request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return photos, fmt.Errorf("unsuccessful request to image GET API: %d", response.StatusCode)
	}

	responseBody, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(responseBody, &photos)
	if err != nil {
		return photos, fmt.Errorf("error unmarshalling image GET response: %w", err)
	}

	return photos, nil
}

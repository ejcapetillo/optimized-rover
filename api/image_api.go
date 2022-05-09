package api

import "errors"

type ImageAPI interface {
	GetImages(nasaUrl string) error
}

type imageAPI struct {
}

func NewImageAPI() ImageAPI {
	return &imageAPI{}
}

func (api *imageAPI) GetImages(nasaUrl string) error {
	if nasaUrl == "" {
		return errors.New("missing URL")
	}
	return nil
}

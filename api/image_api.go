package api

type ImageAPI interface {
	GetImages() error
}

type imageAPI struct {
}

func NewImageAPI() ImageAPI {
	return &imageAPI{}
}

func (api *imageAPI) GetImages() error {
	return nil
}

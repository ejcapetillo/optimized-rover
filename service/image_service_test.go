package service

import (
	"errors"
	"fmt"
	"github.com/ejcapetillo/optimized-rover/mocks"
	"github.com/ejcapetillo/optimized-rover/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetImagesSingleThread(t *testing.T) {
	// Inject mocks for the API tier
	api := new(mocks.ImageAPI)
	service := NewImageService(api)

	// Setup mock expectations
	numDaysToCompare = 3
	dateList := initializeDateList()
	nasaURL1 := fmt.Sprintf("%s/%s/photos?api_key=%s&earth_date=%s", nasaRoverAPIRoot, model.Curiosity, model.NasaAPIKey, dateList[2])
	nasaURL2 := fmt.Sprintf("%s/%s/photos?api_key=%s&earth_date=%s", nasaRoverAPIRoot, model.Curiosity, model.NasaAPIKey, dateList[1])
	nasaURL3 := fmt.Sprintf("%s/%s/photos?api_key=%s&earth_date=%s", nasaRoverAPIRoot, model.Curiosity, model.NasaAPIKey, dateList[0])
	day1Images := &model.PhotoWrapper{}
	day2Images := &model.PhotoWrapper{Photos: []*model.Photo{{Id: 1}, {Id: 2}}}
	day3Images := &model.PhotoWrapper{Photos: []*model.Photo{{Id: 3}, {Id: 4}, {Id: 5}}}
	day1Expectation := model.DailyPhoto{EarthDate: dateList[2], Count: 0}
	day2Expectation := model.DailyPhoto{EarthDate: dateList[1], Count: 2}
	day3Expectation := model.DailyPhoto{EarthDate: dateList[0], Count: 3}

	// Setup Mock Results
	api.On("GetImages", mock.Anything).Return(nil, errors.New("mock error message")).Once()
	api.On("GetImages", nasaURL3).Return(day3Images, nil).Once()
	api.On("GetImages", nasaURL2).Return(day2Images, nil).Once()
	api.On("GetImages", nasaURL1).Return(day1Images, nil).Once()

	// Test Case 1: Error thrown inside the API tier
	result, err := service.GetImagesSingleThread()
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("mock error message"), err)

	// Test Case 2: Ensure that the method is aggregating and collecting correctly
	result, err = service.GetImagesSingleThread()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, result, day1Expectation)
	assert.Contains(t, result, day2Expectation)
	assert.Contains(t, result, day3Expectation)
}

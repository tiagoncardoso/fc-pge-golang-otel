package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/test/mocks"
	"net/http"
	"testing"
)

func Test_GivenAValidZipCode_WhenRequestZipData_ThenReturnZipData(t *testing.T) {
	mockHttpRequest := &mocks.HttpRequestMock{}
	zipCode := "74333110"
	ctx := context.Background()
	var req *http.Request

	usecase := NewRequestWeather("http://localhost:8081/temperature/{ZIP}")
	successReturn := dto.WeatherDetailsOutputDto{
		TempK: 298.25,
		TempF: 77.18,
		TempC: 25.1,
		City:  "Goiânia",
	}

	mockHttpRequest.On("HttpGetRequest").Return(successReturn, nil)
	zipData, err := usecase.Execute(zipCode, ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, zipData)
	assert.NotEmpty(t, zipData.City)
	assert.Equal(t, successReturn.City, zipData.City)
	assert.Equal(t, successReturn.TempC, zipData.TempC)
	assert.Equal(t, successReturn.TempF, zipData.TempF)
	assert.Equal(t, successReturn.TempK, zipData.TempK)
}

func Test_GivenAnInvalidZipCode_WhenRequestZipData_ThenReturnError(t *testing.T) {
	mockHttpRequest := &mocks.HttpRequestMock{}
	zipCode := "7433311"
	ctx := context.Background()
	var req *http.Request

	usecase := NewRequestWeather("http://localhost:8081/temperature/{ZIP}")

	mockHttpRequest.On("HttpGetRequest").Return(dto.WeatherDetailsOutputDto{}, nil)
	zipData, err := usecase.Execute(zipCode, ctx, req)

	assert.Error(t, err)
	assert.NotNil(t, zipData)
	assert.Empty(t, zipData.City)
}

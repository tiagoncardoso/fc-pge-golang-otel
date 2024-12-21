package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/test/mocks"
	"go.opentelemetry.io/otel/trace/noop"
	"testing"
)

func Test_GivenAValidCityName_WhenRequestWeatherData_ThenReturnWeatherData(t *testing.T) {
	mockHttpRequest := &mocks.HttpRequestMock{}
	cityName := "Goiânia"
	tracer := noop.Tracer{}
	ctx := context.Background()

	usecase := NewRequestWeatherData("https://api.openweathermap.org/data/2.5/weather?q={CITY}&appid=123456789", tracer)

	mockHttpRequest.On("HttpGetRequest").Return(dto.WeatherApiDto{}, nil)
	weatherData, err := usecase.Execute(cityName, ctx)

	assert.NoError(t, err)
	assert.NotNil(t, weatherData)
}

func Test_GivenAnEmptyCityName_WhenRequestWeatherData_ThenReturnError(t *testing.T) {
	mockHttpRequest := &mocks.HttpRequestMock{}
	cityName := ""
	tracer := noop.Tracer{}
	ctx := context.Background()

	usecase := NewRequestWeatherData("https://api.openweathermap.org/data/2.5/weather?q={CITY}&appid=123456789", tracer)

	mockHttpRequest.On("HttpGetRequest").Return(dto.WeatherApiDto{}, nil)
	weatherData, err := usecase.Execute(cityName, ctx)

	assert.Error(t, err)
	assert.NotNil(t, weatherData)
	assert.Equal(t, err.Error(), "city name is empty")
}

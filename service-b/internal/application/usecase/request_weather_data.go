package usecase

import (
	"errors"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/pkg/http_request"
	"net/url"
	"strings"
)

type RequestWeatherData struct {
	weatherDataApiUrl string
}

func NewRequestWeatherData(weatherDataApiUrl string) *RequestWeatherData {
	return &RequestWeatherData{
		weatherDataApiUrl: weatherDataApiUrl,
	}
}

func (r *RequestWeatherData) Execute(cityName string) (dto.WeatherApiDto, error) {
	if cityName == "" {
		return dto.WeatherApiDto{}, errors.New("city name is empty")
	}

	weatherUrl := makeWeatherApiUrl(r.weatherDataApiUrl, cityName)
	weatherData, err := http_request.HttpGetRequest[dto.WeatherApiDto](weatherUrl)

	if err != nil {
		return dto.WeatherApiDto{}, err
	}

	return weatherData, nil
}

func makeWeatherApiUrl(weatherBaseUrl string, cityName string) string {
	return strings.Replace(weatherBaseUrl, "{CITY}", url.QueryEscape(cityName), 1)
}

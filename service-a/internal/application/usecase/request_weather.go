package usecase

import (
	"fmt"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/pkg/http_request"
	"strings"
)

type RequestWeather struct {
	ApiUrl string
}

func NewRequestWeather(apiUrl string) *RequestWeather {
	return &RequestWeather{
		ApiUrl: apiUrl,
	}
}

func (r *RequestWeather) Execute(zipCode string) (dto.WeatherDetailsOutputDto, error) {
	url := makeZipApiUrl(r.ApiUrl, zipCode)
	fmt.Println(url)
	respWeather, err := http_request.HttpGetRequest[dto.WeatherDetailsOutputDto](url)
	if err != nil {
		return dto.WeatherDetailsOutputDto{}, err
	}

	return respWeather, nil
}

func makeZipApiUrl(apiBaseUrl string, zipCode string) string {
	return strings.Replace(apiBaseUrl, "{ZIP}", zipCode, 1)
}

package usecase

import (
	"context"
	"errors"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/pkg/http_request"
	"go.opentelemetry.io/otel/trace"
	"net/url"
	"strings"
)

type RequestWeatherData struct {
	weatherDataApiUrl string
	Tracer            trace.Tracer
}

func NewRequestWeatherData(weatherDataApiUrl string, tracer trace.Tracer) *RequestWeatherData {
	return &RequestWeatherData{
		weatherDataApiUrl: weatherDataApiUrl,
		Tracer:            tracer,
	}
}

func (r *RequestWeatherData) Execute(cityName string, ctx context.Context) (dto.WeatherApiDto, error) {
	ctx, span := r.Tracer.Start(ctx, "Request.WeatherAPI")
	defer span.End()

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

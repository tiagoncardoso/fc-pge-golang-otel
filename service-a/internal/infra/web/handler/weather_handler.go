package handler

import (
	"encoding/json"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/helper"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type WeatherHandler struct {
	RequestWeatherUsecase *usecase.RequestWeather
	Tracer                trace.Tracer
	ServiceNameRequest    string
}

func NewWeatherHandler(requestWeatherApiUsecase *usecase.RequestWeather, tracer trace.Tracer, serviceNameRequest string) *WeatherHandler {
	return &WeatherHandler{
		RequestWeatherUsecase: requestWeatherApiUsecase,
		Tracer:                tracer,
		ServiceNameRequest:    serviceNameRequest,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, span := h.Tracer.Start(ctx, h.ServiceNameRequest)
	defer span.End()

	var requestBody struct {
		Cep string `json:"cep"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	if !helper.IsValidZipCode(requestBody.Cep) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))

		return
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))
	weatherData, err := h.RequestWeatherUsecase.Execute(requestBody.Cep, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))

		return
	}

	output := dto.WeatherDetailsOutputDto{
		City:  weatherData.City,
		TempC: weatherData.TempC,
		TempF: weatherData.TempF,
		TempK: weatherData.TempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package web

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/helper"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/usecase"
	"net/http"
)

type WeatherHandler struct {
	ZipApiUsecase     *usecase.RequestZipData
	WeatherApiUsecase *usecase.RequestWeatherData
}

func NewWeatherHandler(zipApiUsecase *usecase.RequestZipData, weatherApiUsecase *usecase.RequestWeatherData) *WeatherHandler {
	return &WeatherHandler{
		ZipApiUsecase:     zipApiUsecase,
		WeatherApiUsecase: weatherApiUsecase,
	}
}

func (h *WeatherHandler) GetWeatherByZip(w http.ResponseWriter, r *http.Request) {
	var zipCode = chi.URLParam(r, "cep")

	if !helper.IsValidZipCode(zipCode) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))

		return
	}

	zipData, err := h.ZipApiUsecase.Execute(helper.SanitizeZipCode(zipCode))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))

		return
	}

	if zipData.Erro == "true" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))

		return
	}

	weatherData, err := h.WeatherApiUsecase.Execute(zipData.Localidade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output := dto.WeatherDetailsOutputDto{
		City:  zipData.Localidade,
		TempC: weatherData.Current.TempC,
		TempF: helper.ConvertCelsiusToFarenheig(weatherData.Current.TempC),
		TempK: helper.ConvertCelsiusToKelvin(weatherData.Current.TempC),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

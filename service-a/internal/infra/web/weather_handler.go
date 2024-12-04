package web

import (
	"encoding/json"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/helper"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/usecase"
	"net/http"
)

type WeatherHandler struct {
	RequestWeatherUsecase *usecase.RequestWeather
}

func NewWeatherHandler(requestWeatherApiUsecase *usecase.RequestWeather) *WeatherHandler {
	return &WeatherHandler{
		RequestWeatherUsecase: requestWeatherApiUsecase,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
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

	weatherData, err := h.RequestWeatherUsecase.Execute(requestBody.Cep)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))

		return
	}

	// TODO: >> Olhar como tratar o erro de resposta da api b
	//if zipData.Erro == "true" {
	//	w.WriteHeader(http.StatusNotFound)
	//	w.Write([]byte("can not find zipcode"))
	//
	//	return
	//}

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

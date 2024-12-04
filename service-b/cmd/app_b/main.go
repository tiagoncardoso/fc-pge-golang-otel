package main

import (
	"fmt"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/config"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/usecase"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/infra/web"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/infra/web/webserver"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	findZip := usecase.NewRequestZipData(conf.ApiUrlZip)
	findWeather := usecase.NewRequestWeatherData(conf.ApiUrlWeather + "" + conf.ApiKeyWeather)

	webServer := webserver.NewWebServer(conf.WebServerPort)
	zipWeatherHandler := web.NewWeatherHandler(findZip, findWeather)

	webServer.AddHandler("/temperature/{cep}", "GET", zipWeatherHandler.GetWeatherByZip)

	fmt.Println("Starting web server on port", conf.WebServerPort)
	webServer.Start()
}

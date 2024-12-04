package main

import (
	"fmt"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/config"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/usecase"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/infra/web"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/infra/web/webserver"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	findDataUsecase := usecase.NewRequestWeather(conf.ApiService)

	webServer := webserver.NewWebServer(conf.WebServerPort)
	weatherHandler := web.NewWeatherHandler(findDataUsecase)

	webServer.AddHandler("/", "POST", weatherHandler.GetWeather)

	fmt.Println("Starting web server on port", conf.WebServerPort)
	webServer.Start()
}

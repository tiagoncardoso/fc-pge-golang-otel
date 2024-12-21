package main

import (
	"context"
	"fmt"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/config"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/usecase"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/infra/web"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/infra/web/handler"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/pkg/opentelemetry"
	"go.opentelemetry.io/otel"
	"log"
	"os"
	"os/signal"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ot := opentelemetry.NewOpenTelemetry(conf.ServiceName, conf.CollectorUrl)
	otp, err := ot.InitProvider()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := otp(ctx); err != nil {
			log.Fatalln("failed to shutdown TraceProvider", err)
		}
	}()

	tracer := otel.Tracer(conf.ServiceNameRequest)

	findZip := usecase.NewRequestZipData(conf.ApiUrlZip, tracer)
	findWeather := usecase.NewRequestWeatherData(conf.ApiUrlWeather+""+conf.ApiKeyWeather, tracer)

	webServer := web.NewWebServer(conf.WebServerPort)
	zipWeatherHandler := handler.NewWeatherHandler(findZip, findWeather, tracer, conf.ServiceName)

	webServer.AddHandler("/temperature/{cep}", "GET", zipWeatherHandler.GetWeatherByZip)

	fmt.Println("Starting web server on port", conf.WebServerPort)
	webServer.Start()
}

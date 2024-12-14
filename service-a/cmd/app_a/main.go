package main

import (
	"context"
	"fmt"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/config"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/application/usecase"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/infra/web"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/internal/infra/web/handler"
	"github.com/tiagoncardoso/fc-pge-golang-otel-a/pkg/opentelemetry"
	"go.opentelemetry.io/otel"
	"log"
	"os"
	"os/signal"
	"time"
)

//func init() {
//	viper.AutomaticEnv()
//}

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

	tracer := otel.Tracer("fullcycle-service-a")

	findDataUsecase := usecase.NewRequestWeather(conf.ApiService)

	webServer := web.NewWebServer(conf.WebServerPort)
	weatherHandler := handler.NewWeatherHandler(findDataUsecase, tracer, conf.ServiceNameRequest)

	webServer.AddHandler("/", "POST", weatherHandler.GetWeather)

	go func() {
		fmt.Println("Starting web server on port", conf.WebServerPort)
		webServer.Start()
	}()

	select {
	case <-sigCh:
		log.Println("Gracefully shutting down")
	case <-ctx.Done():
		log.Println("Context done")
	}

	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	//conf, err := config.LoadConfig()
	//if err != nil {
	//	panic(err)
	//}
	//
	//findDataUsecase := usecase.NewRequestWeather(conf.ApiService)
	//
	//webServer := webserver.NewWebServer(conf.WebServerPort)
	//weatherHandler := web.NewWeatherHandler(findDataUsecase)
	//
	//webServer.AddHandler("/", "POST", weatherHandler.GetWeather)
	//
	//fmt.Println("Starting web server on port", conf.WebServerPort)
	//webServer.Start()
}

FROM golang:1.23 as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather-app ./cmd/weather_zip_app/main.go

FROM golang:1.23 as prod

WORKDIR /app
COPY --from=build /app/weather-app .
COPY --from=build /app/.env .env

EXPOSE 8080

ENTRYPOINT ["./weather-app"]
FROM golang:1.23 AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service-b ./cmd/app_b/main.go

FROM golang:1.23 AS prod

WORKDIR /app
COPY --from=build /app/service-b .
COPY --from=build /app/.env .env

EXPOSE 8080

ENTRYPOINT ["./service-b"]
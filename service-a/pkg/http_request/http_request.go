package http_request

import (
	"context"
	"encoding/json"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
)

type RequestData interface {
	interface{}
}

func HttpGetRequest[T RequestData](url string, ctx context.Context, r *http.Request) (T, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	return httpRequest[T](url, http.MethodGet, headers, ctx, r)
}

func httpRequest[T RequestData](url string, method string, headers map[string]string, ctx context.Context, r *http.Request) (T, error) {
	var apiResponse T

	cl := http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return apiResponse, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := cl.Do(req)
	if err != nil {
		return apiResponse, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return apiResponse, err
	}

	return apiResponse, nil
}

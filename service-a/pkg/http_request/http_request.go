package http_request

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

const (
	API_REQUEST_TIMEOUT = 5
)

type RequestData interface {
	interface{}
}

func HttpGetRequest[T RequestData](url string) (T, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	return httpRequest[T](url, http.MethodGet, headers)
}

func httpRequest[T RequestData](url string, method string, headers map[string]string) (T, error) {
	var apiResponse T

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(API_REQUEST_TIMEOUT))
	defer cancel()

	cl := http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return apiResponse, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

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

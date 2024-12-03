package usecase

import (
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/helper"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/pkg/http_request"
	"strings"
)

type RequestZipData struct {
	zipDataApiUrl string
}

func NewRequestZipData(zipDataApiUrl string) *RequestZipData {
	return &RequestZipData{
		zipDataApiUrl: zipDataApiUrl,
	}
}

func (r *RequestZipData) Execute(zipCode string) (dto.ViaCepApiDto, error) {
	zipUrl := makeZipApiUrl(r.zipDataApiUrl, zipCode)
	zipData, err := http_request.HttpGetRequest[dto.ViaCepApiDto](zipUrl)
	if err != nil {
		return dto.ViaCepApiDto{}, err
	}

	return zipData, nil
}

func makeZipApiUrl(zipCodeBaseUrl string, zipCode string) string {
	strZipCode := helper.SanitizeZipCode(zipCode)

	return strings.Replace(zipCodeBaseUrl, "{ZIP}", strZipCode, 1)
}

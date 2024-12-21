package usecase

import (
	"context"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/helper"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/pkg/http_request"
	"go.opentelemetry.io/otel/trace"
	"strings"
)

type RequestZipData struct {
	zipDataApiUrl string
	Tracer        trace.Tracer
}

func NewRequestZipData(zipDataApiUrl string, tracer trace.Tracer) *RequestZipData {
	return &RequestZipData{
		zipDataApiUrl: zipDataApiUrl,
		Tracer:        tracer,
	}
}

func (r *RequestZipData) Execute(zipCode string, ctx context.Context) (dto.ViaCepApiDto, error) {
	ctx, span := r.Tracer.Start(ctx, "Request.ViaCEP")
	defer span.End()

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

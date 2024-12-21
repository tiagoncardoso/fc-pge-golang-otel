package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/test/mocks"
	"go.opentelemetry.io/otel/trace/noop"
	"testing"
)

func Test_GivenAValidZipCode_WhenRequestZipData_ThenReturnZipData(t *testing.T) {
	mockHttpRequest := &mocks.HttpRequestMock{}
	zipCode := "74333-110"
	ctx := context.Background()
	tracer := noop.Tracer{}

	usecase := NewRequestZipData("https://viacep.com.br/ws/{ZIP}/json/", tracer)
	successReturn := dto.ViaCepApiDto{
		Cep: zipCode,
	}

	mockHttpRequest.On("HttpGetRequest").Return(successReturn, nil)
	zipData, err := usecase.Execute(zipCode, ctx)

	assert.NoError(t, err)
	assert.NotNil(t, zipData)
	assert.Equal(t, zipCode, zipData.Cep)
}

func Test_GivenAnInvalidZipCode_WhenRequestZipData_ThenReturnError(t *testing.T) {
	mockHttpRequest := &mocks.HttpRequestMock{}
	zipCode := "7433311"
	tracer := noop.Tracer{}
	ctx := context.Background()

	usecase := NewRequestZipData("https://viacep.com.br/ws/{ZIP}/json/", tracer)

	mockHttpRequest.On("HttpGetRequest").Return(dto.ViaCepApiDto{}, nil)
	zipData, err := usecase.Execute(zipCode, ctx)

	assert.Error(t, err)
	assert.NotNil(t, zipData)
	assert.Empty(t, zipData.Cep)
}

package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/test/mocks"
	"testing"
)

func Test_GivenAValidZipCode_WhenRequestZipData_ThenReturnZipData(t *testing.T) {
	mockHttpRequest := &mocks.HttpRequestMock{}
	zipCode := "74333-110"
	usecase := NewRequestZipData("https://viacep.com.br/ws/{ZIP}/json/")
	successReturn := dto.ViaCepApiDto{
		Cep: zipCode,
	}

	mockHttpRequest.On("HttpGetRequest").Return(successReturn, nil)
	zipData, err := usecase.Execute(zipCode)

	assert.NoError(t, err)
	assert.NotNil(t, zipData)
	assert.Equal(t, zipCode, zipData.Cep)
}

func Test_GivenAnInvalidZipCode_WhenRequestZipData_ThenReturnError(t *testing.T) {
	mockHttpRequest := &mocks.HttpRequestMock{}
	zipCode := "7433311"
	usecase := NewRequestZipData("https://viacep.com.br/ws/{ZIP}/json/")

	mockHttpRequest.On("HttpGetRequest").Return(dto.ViaCepApiDto{}, nil)
	zipData, err := usecase.Execute(zipCode)

	assert.Error(t, err)
	assert.NotNil(t, zipData)
	assert.Empty(t, zipData.Cep)
}

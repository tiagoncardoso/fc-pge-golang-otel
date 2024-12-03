package e2e

import (
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/application/usecase"
	"github.com/tiagoncardoso/fc-pge-golang-otel-b/internal/infra/web"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RequestWeatherDataTestSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *RequestWeatherDataTestSuite) SetupSuite() {
	ApiUrlZip := "http://viacep.com.br/ws/{ZIP}/json/"
	ApiUrlWeather := "https://api.weatherapi.com/v1/current.json?q={CITY}&key="
	ApiKeyWeather := "b0e4e64ea8d9434bac4142752241511"

	zipCodeUsecase := usecase.NewRequestZipData(ApiUrlZip)
	weatherDataUsecase := usecase.NewRequestWeatherData(ApiUrlWeather + "" + ApiKeyWeather)
	zipWeatherHandler := web.NewWeatherHandler(zipCodeUsecase, weatherDataUsecase)

	r := chi.NewRouter()
	r.Get("/temperature/{cep}", zipWeatherHandler.GetWeatherByZip)

	s.server = httptest.NewServer(r)
}

func (s *RequestWeatherDataTestSuite) TearDownSuite() {
	s.server.Close()
}

func TestRequestWeatherDataTestSuite(t *testing.T) {
	suite.Run(t, new(RequestWeatherDataTestSuite))
}

func (s *RequestWeatherDataTestSuite) Test_GivenAValidAndExistsZipCode_WhenRequestWeatherData_ThenReturnWeatherData() {
	validZipCode := "74333-110"

	resp, err := http.Get(s.server.URL + "/temperature/" + validZipCode)
	s.NoError(err)

	s.Equal(http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	s.NoError(err)
	s.Contains(string(body), "temp_C")
	s.Contains(string(body), "temp_F")
	s.Contains(string(body), "temp_K")
}

func (s *RequestWeatherDataTestSuite) Test_GivenAnInvalidZipCode_WhenRequestWeatherData_ThenReturn422Error() {
	invalidZipCode := "12312-3"

	resp, err := http.Get(s.server.URL + "/temperature/" + invalidZipCode)
	s.NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusUnprocessableEntity, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	s.NoError(err)
	s.Equal("invalid zipcode", string(body))
}

func (s *RequestWeatherDataTestSuite) Test_GivenAValidZipCode_WhenRequestNotFoundWeatherData_ThenReturn404Error() {
	notFoundZipCode := "70000000"

	resp, err := http.Get(s.server.URL + "/temperature/" + notFoundZipCode)
	s.NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusNotFound, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	s.NoError(err)
	s.Equal("can not find zipcode", string(body))
}

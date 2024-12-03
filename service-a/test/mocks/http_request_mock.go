package mocks

import "github.com/stretchr/testify/mock"

type HttpRequestMock struct {
	mock.Mock
}

func (m *HttpRequestMock) HttpGetRequest(url string) (interface{}, error) {
	args := m.Called(url)
	return args.Get(0), args.Error(1)
}

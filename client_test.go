package shopstyle_test

import (
	"github.com/cmethvin/shopstyle"
	"net/http"
)

type MockHttpClient struct {
	DoCalled bool
	DoFunc   func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	m.DoCalled = true
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return nil, nil
}

var _ shopstyle.HttpClient = &MockHttpClient{}

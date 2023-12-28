package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	expectedCode := 200
	req, err := http.NewRequest("GET", "/cafe", nil) // здесь нужно создать запрос к сервису
	if err != nil {
		require.NoError(t, err)
	}
	query := req.URL.Query()
	query.Add("count", "10")
	query.Add("city", "moscow")

	req.URL.RawQuery = query.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	resultBody := strings.Split(responseRecorder.Body.String(), ",")
	resultCode := responseRecorder.Code

	require.Equal(t, resultCode, expectedCode)
	require.Equal(t, totalCount, len(resultBody))
}

func TestWrongCity(t *testing.T) {
	expectedError := "wrong city value"
	expectedCode := 400
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}

	query := req.URL.Query()
	query.Add("count", "4")
	query.Add("city", "omsk")
	req.URL.RawQuery = query.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	resultCode := responseRecorder.Code
	require.Equal(t, expectedCode, resultCode)
	result := responseRecorder.Body.String()
	require.Equal(t, result, expectedError)
}

func TestCorrectQuery(t *testing.T) {
	expectedCode := 200
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}

	query := req.URL.Query()
	query.Add("count", "3")
	query.Add("city", "moscow")
	req.URL.RawQuery = query.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	resultCode := responseRecorder.Code
	require.Equal(t, expectedCode, resultCode)
	require.NotEmpty(t, responseRecorder.Body)
}

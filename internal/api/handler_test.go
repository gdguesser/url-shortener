package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockURLShortenerService struct {
	mock.Mock
}

func (m *MockURLShortenerService) Shorten(longURL string) (string, error) {
	args := m.Called(longURL)
	return args.String(0), args.Error(1)
}

func (m *MockURLShortenerService) Resolve(shortCode string) (string, error) {
	args := m.Called(shortCode)
	return args.String(0), args.Error(1)
}

func TestShortenURLHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	router := gin.Default()
	router.POST("/shorten", handler.ShortenURL)

	longURL := "https://www.example.com"
	shortCode := "abc123"
	mockService.On("Shorten", longURL).Return(shortCode, nil)

	requestBody, _ := json.Marshal(gin.H{"long_url": longURL})
	req, _ := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), shortCode)
}

func TestResolveURLHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	router := gin.Default()
	router.GET("/:short_code", handler.ResolveURL)

	shortCode := "abc123"
	longURL := "https://www.example.com"
	mockService.On("Resolve", shortCode).Return(longURL, nil)

	req, _ := http.NewRequest(http.MethodGet, "/abc123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, longURL, w.Header().Get("Location"))
}

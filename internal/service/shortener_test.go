package service

import (
	"testing"

	"github.com/gdguesser/url-shortener/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockURLRepository struct {
	mock.Mock
}

func (m *MockURLRepository) Create(url *models.URL) error {
	args := m.Called(url)
	return args.Error(0)
}

func (m *MockURLRepository) GetByShortCode(code string) (*models.URL, error) {
	args := m.Called(code)
	return args.Get(0).(*models.URL), args.Error(1)
}

func (m *MockURLRepository) IncrementCounter(code string) error {
	args := m.Called(code)
	return args.Error(0)
}

func TestShorten(t *testing.T) {
	mockRepo := new(MockURLRepository)
	service := NewURLShortenerService(mockRepo)

	longURL := "https://example.com"
	mockRepo.On("Create", mock.Anything).Return(nil)

	_, err := service.Shorten(longURL)

	assert.NoError(t, err)
}

func TestResolve(t *testing.T) {
	mockRepo := new(MockURLRepository)
	service := NewURLShortenerService(mockRepo)

	shortCode := "abc123"
	longURL := "https://example.com"
	mockRepo.On("IncrementCounter", shortCode).Return(nil)
	mockRepo.On("GetByShortCode", shortCode).Return(&models.URL{LongURL: longURL}, nil)

	result, err := service.Resolve(shortCode)

	assert.NoError(t, err)
	assert.Equal(t, longURL, result)
}

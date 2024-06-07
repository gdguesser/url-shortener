package service

import (
	"net/url"

	"github.com/gdguesser/url-shortener/internal/repository"
	"github.com/gdguesser/url-shortener/pkg/models"
	"github.com/teris-io/shortid"
)

type URLShortenerService interface {
	Shorten(longURL string) (string, error)
	Resolve(shortCode string) (string, error)
}

type urlShortenerService struct {
	repo repository.URLRepository
}

func NewURLShortenerService(repo repository.URLRepository) URLShortenerService {
	return &urlShortenerService{repo: repo}
}

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	return u.String(), nil
}

func (s *urlShortenerService) Shorten(longURL string) (string, error) {
	normalizedURL, err := normalizeURL(longURL)
	if err != nil {
		return "", err
	}

	shortCode, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	url := &models.URL{
		LongURL:   normalizedURL,
		ShortCode: shortCode,
	}

	err = s.repo.Create(url)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *urlShortenerService) Resolve(shortCode string) (string, error) {
	err := s.repo.IncrementCounter(shortCode)
	if err != nil {
		return "", err
	}

	url, err := s.repo.GetByShortCode(shortCode)
	if err != nil {
		return "", err
	}

	return url.LongURL, nil
}

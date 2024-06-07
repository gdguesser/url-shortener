package service

import (
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

func NewUrlShortenerService(repo repository.URLRepository) URLShortenerService {
	return &urlShortenerService{repo: repo}
}

func (s *urlShortenerService) Shorten(longURL string) (string, error) {
	shortCode, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	url := &models.URL{
		LongURL:   longURL,
		ShortCode: shortCode,
	}

	err = s.repo.Create(url)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *urlShortenerService) Resolve(shortCode string) (string, error) {
	url, err := s.repo.GetByShortCode(shortCode)
	if err != nil {
		return "", err
	}

	return url.LongURL, nil
}

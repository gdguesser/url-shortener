package repository

import (
	"github.com/gdguesser/url-shortener/pkg/models"
	"gorm.io/gorm"
)

type URLRepository interface {
	Create(url *models.URL) error
	GetByShortCode(code string) (*models.URL, error)
}

type urlRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{db: db}
}

func (r *urlRepository) Create(url *models.URL) error {
	return r.db.Create(url).Error
}

func (r *urlRepository) GetByShortCode(code string) (*models.URL, error) {
	var url models.URL
	err := r.db.Where("short_code = ?", code).First(&url).Error
	return &url, err
}

package repository

import (
	"testing"

	"github.com/gdguesser/url-shortener/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the test database")
	}

	db.AutoMigrate(&models.URL{})

	cleanup := func() {
		db.Exec("DROP TABLE urls")
	}

	return db, cleanup
}

func TestCreateURL(t *testing.T) {
	db, cleanup := setupTestDB()
	defer cleanup()

	repo := NewURLRepository(db)

	url := &models.URL{
		LongURL:   "https://www.example.com",
		ShortCode: "abc123",
		Counter:   0,
	}

	err := repo.Create(url)
	assert.NoError(t, err)

	var retrievedURL models.URL
	err = db.First(&retrievedURL, "short_code = ?", "abc123").Error
	assert.NoError(t, err)
	assert.Equal(t, url.LongURL, retrievedURL.LongURL)
	assert.Equal(t, url.ShortCode, retrievedURL.ShortCode)
	assert.Equal(t, url.Counter, retrievedURL.Counter)
}

func TestGetByShortCode(t *testing.T) {
	db, cleanup := setupTestDB()
	defer cleanup()

	repo := NewURLRepository(db)

	url := &models.URL{
		LongURL:   "https://www.example.com",
		ShortCode: "abc123",
		Counter:   0,
	}

	db.Create(url)

	retrievedURL, err := repo.GetByShortCode("abc123")
	assert.NoError(t, err)
	assert.Equal(t, url.LongURL, retrievedURL.LongURL)
	assert.Equal(t, url.ShortCode, retrievedURL.ShortCode)
	assert.Equal(t, url.Counter, retrievedURL.Counter)
}

func TestIncrementCounter(t *testing.T) {
	db, cleanup := setupTestDB()
	defer cleanup()

	repo := NewURLRepository(db)

	url := &models.URL{
		LongURL:   "https://www.example.com",
		ShortCode: "abc123",
		Counter:   0,
	}

	db.Create(url)

	err := repo.IncrementCounter("abc123")
	assert.NoError(t, err)

	var retrievedURL models.URL
	err = db.First(&retrievedURL, "short_code = ?", "abc123").Error
	assert.NoError(t, err)
	assert.Equal(t, 1, retrievedURL.Counter)
}

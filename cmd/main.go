package main

import (
	"github.com/gdguesser/url-shortener/internal/api"
	"github.com/gdguesser/url-shortener/internal/repository"
	"github.com/gdguesser/url-shortener/internal/service"
	"github.com/gdguesser/url-shortener/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.URL{})

	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLShortenerService(urlRepo)
	urlHandler := api.NewURLHandler(urlService)

	r := gin.Default()

	r.POST("/shorten", urlHandler.ShortenURL)
	r.GET("/:short_code", urlHandler.ResolveURL)

	r.Run(":8080")
}

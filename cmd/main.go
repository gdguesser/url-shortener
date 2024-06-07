package main

import (
	"html/template"
	"net/http"

	"github.com/gdguesser/url-shortener/internal/api"
	"github.com/gdguesser/url-shortener/internal/repository"
	"github.com/gdguesser/url-shortener/internal/service"
	"github.com/gdguesser/url-shortener/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

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
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		tmpl.ExecuteTemplate(c.Writer, "index.html", nil)
	})

	r.POST("/shorten", func(c *gin.Context) {
		longURL := c.PostForm("long_url")
		shortCode, err := urlService.Shorten(longURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tmpl.ExecuteTemplate(c.Writer, "result.html", gin.H{"ShortCode": shortCode})
	})

	r.GET("/:short_code", urlHandler.ResolveURL)

	r.Run(":8080")
}

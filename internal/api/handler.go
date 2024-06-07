package api

import (
	"net/http"

	"github.com/gdguesser/url-shortener/internal/service"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service service.URLShortenerService
}

func NewURLHandler(service service.URLShortenerService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) ShortenURL(c *gin.Context) {
	var request struct {
		LongURL string `json:"long_url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortCode, err := h.service.Shorten(request.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_code": shortCode})
}

func (h *URLHandler) ResulveURL(c *gin.Context) {
	shortCode := c.Param("short_code")

	longURL, err := h.service.Resolve(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, longURL)
}

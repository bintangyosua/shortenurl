package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/bintangyosua/shortenurl/config"
	"github.com/bintangyosua/shortenurl/models"
	"github.com/bintangyosua/shortenurl/views"

	"github.com/labstack/echo/v4"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	s := make([]byte, length)
	for i := range s {
		s[i] = charset[rand.Intn(len(charset))]
	}
	return string(s)
}

func Index(c echo.Context) error {
	var urls []models.URL

	result := config.DB.Order("created_at DESC").Find(&urls)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch URLs")
	}

	// Kirim data ke view
	return views.Index(urls).Render(c.Request().Context(), c.Response().Writer)
}

func ShortenURL(c echo.Context) error {
	original := c.FormValue("original_url")
	shortCode := randomString(6)

	url := models.URL{
		OriginalURL: original,
		ShortCode:   shortCode,
	}

	config.DB.Create(&url)

	var urls []models.URL

	result := config.DB.Order("created_at DESC").Find(&urls)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch URLs")
	}

	return views.ListUrl(urls).Render(c.Request().Context(), c.Response().Writer)
}

func Redirect(c echo.Context) error {
	code := c.Param("code")
	var url models.URL
	result := config.DB.Where("short_code = ?", code).First(&url)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "URL not found")
	}
	return c.Redirect(http.StatusFound, url.OriginalURL)
}


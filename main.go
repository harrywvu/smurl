package main

import (
	// "fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type URL struct {
	OriginalURL string `json:"original_url"`
}

func shortenURL(c *gin.Context) {
	var req URL

	err := c.ShouldBindJSON(&req)
	if err != nil {
		respondError(c, 400, "invalid_json", "invalid request")
		return
	}

	trimmed := strings.TrimSpace(req.OriginalURL)
	// validation
	if trimmed == "" {
		respondError(c, 400, "invalid_json", "original_url is required")
		return
	}

	if len(trimmed) > 2048 {
		respondError(c, 400, "invalid_url", "url is too long")
		return
	}

	u, err := url.Parse(trimmed)
	if err != nil {
		respondError(c, 400, "invalid_url", "url can't be parsed")
		return
	}

	lowerscheme := strings.ToLower(u.Scheme)
	if lowerscheme != "http" && lowerscheme != "https" {
		respondError(c, 400, "invalid_url", "url isn't a valid http/https url")
		return
	}

	if len(u.Host) == 0 {
		respondError(c, 400, "invalid_url", "url host can't be empty")
		return
	}

	created := time.Now().UTC().Format(time.RFC3339)

	// placeholder values for now
	c.JSON(201, gin.H{
		"code":         "abc1234",
		"short_url":    "http://localhost:8080/abc1234",
		"original_url": trimmed,
		"created_at":   created,
	})
}

func main() {
	r := gin.Default()

	// take original url and return the shortened url
	// router.POST("/shorten", shortenURL)

	// router.GET()

	r.POST("/api/v1/urls", shortenURL)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run()
}

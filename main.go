package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type URL struct {
	OriginalURL string `json:"OriginalUrl"`
}

func shortenURL(c *gin.Context) {
	var req URL

	err := c.ShouldBindJSON(&req)
	url := req.OriginalURL

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if req.OriginalURL == "" {
		c.JSON(400, gin.H{"error": "url can't be an empty string"})
		return
	}

	if !IsValidURL(req.OriginalURL) || !IsURLLive(req.OriginalURL) {
		c.JSON(400, gin.H{"error": "url invalid"})
		return
	}

	// pretend we did a process that shortened the url
	url = "short678"

	c.JSON(http.StatusOK, gin.H{"message": url})
}

func main() {
	r := gin.Default()

	// take original url and return the shortened url
	// router.POST("/shorten", shortenURL)

	// router.GET()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run()
}

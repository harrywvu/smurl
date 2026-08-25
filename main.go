package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type URL struct {
	OriginalUrl string `json:"OriginalUrl"`
}

func shortenURL(c *gin.Context) {
	var req URL

	if req.OriginalUrl == "" {
		c.JSON(400, gin.H{"error": "url can't be an empty string"})
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	url := req.OriginalUrl
	// pretend we did a process that shortened the url
	url = "short678"

	c.JSON(http.StatusOK, gin.H{"message": url})
}

func main() {
	router := gin.Default()

	// take original url and return the shortened url
	router.POST("/shorten", shortenURL)

	// router.GET()

	router.Run()
}

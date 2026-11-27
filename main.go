package main

import (
	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type URL struct {
	OriginalURL string `json:"original_url"`
}

func (s *Server) shortenURL(c *gin.Context) {
	
	var req URL

	err := c.ShouldBindJSON(&req)

	if err != nil {
		respondError(c, 400, "invalid_json", "invalid request")
		return
	}

	trimmed, err := validateURL(req.OriginalURL)
	if err != nil {
		respondError(c, 400, "invalid_url", err.Error())
		return
	}

	created := time.Now().UTC().Format(time.RFC3339)

	for {
		code, err := newCode()

		if err != nil {
			respondError(c, 500, "internal_error", "could not generate code")
			return
		}

		if !s.store.Create(code, trimmed) {
			continue
		}

		shortURL := "http://localhost:8080/" + code
	
		c.JSON(201, gin.H{
			"code":         code,
			"short_url":    shortURL,
			"original_url": trimmed,
			"created_at":   created,
		})
		return
	}
}

func newRouter() *gin.Engine {
	r := gin.Default()

	srv := & Server{store: NewStore()}

	r.POST("/api/v1/urls", srv.shortenURL)
	r.GET("/:code", srv.redirect)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}

func main() {
	newRouter().Run()
}

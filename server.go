package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store *Store
}

func (s *Server) redirect(c *gin.Context) {
	code := c.Param("code")

	url, ok := s.store.Resolve(code)
	if !ok {
		respondError(c, 404, "not_found", "short_url not found")
		return
	}

	c.Redirect(http.StatusFound, url)
}
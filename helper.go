package main

import (
	"errors"
	"net/url"
	"strings"

)



func isReserved(code string) bool {
	switch code {
	case "healthz", "metrics":
		return true
	default:
		return false
	}
}

func validateURL(raw string) (string, error){

	trimmed := strings.TrimSpace(raw)
	// validation
	if trimmed == "" {
		return "", errors.New("original_url is required")
	}

	if len(trimmed) > 2048 {
		return "", errors.New("url is too long")
	}

	u, err := url.Parse(trimmed)
	if err != nil {
		return "", errors.New("url can't be parsed")
	}

	lowerscheme := strings.ToLower(u.Scheme)

	if lowerscheme != "http" && lowerscheme != "https" {
		return "", errors.New("url isn't a valid http/https url")
	}

	if len(u.Host) == 0 {
		return "", errors.New("url host can't be empty")
	}

	return trimmed, nil
}
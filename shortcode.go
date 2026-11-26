package main

import (
	"crypto/rand"
	"math/big"
)

const (
	codeLength = 7
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func newCode() (string, error) {
	b := make([]byte, codeLength)
	max := big.NewInt(int64(len(charset)))

	for {
		for i := range b {
			j, err := rand.Int(rand.Reader, max)
			if err != nil {
				return "", err
			}
			b[i] = charset[j.Int64()]
		}

		code := string(b)
		if !isReserved(code) {
			return code, nil
		}
	}
}

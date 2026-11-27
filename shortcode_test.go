package main

import (
	"strings"
	"testing"
)

func TestNewCode(t *testing.T) {
	for i := 0; i < 100; i++ {
		code, err := newCode()
		if err != nil {
			t.Fatalf("newCode() error: %v", err)
		}
		if len(code) != 7 {
			t.Fatalf("newCode() = %q, length %d, want 7", code, len(code))
		}
		for _, char := range code {
			if !strings.ContainsRune(charset, char) {
				t.Fatalf("newCode() = %q, contains invalid character %q", code, char)
			}
		}
	}
}

func TestIsReserved(t *testing.T) {
	for _, code := range []string{"healthz", "metrics"} {
		if !isReserved(code) {
			t.Errorf("isReserved(%q) = false, want true", code)
		}
	}
	if isReserved("abc1234") {
		t.Error("isReserved(\"abc1234\") = true, want false")
	}
}

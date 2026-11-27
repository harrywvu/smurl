package main

import "testing"

func TestStoreCreateAndResolve(t *testing.T) {
	store := NewStore()
	const code = "abc1234"
	const original = "https://example.com/first"

	if !store.Create(code, original) {
		t.Fatal("Create returned false for a new code")
	}
	if store.Create(code, "https://example.com/second") {
		t.Fatal("Create returned true for a duplicate code")
	}
	if got, ok := store.Resolve(code); got != original || !ok {
		t.Errorf("Resolve(%q) = (%q, %t), want (%q, true)", code, got, ok, original)
	}
	if got, ok := store.Resolve("missing"); got != "" || ok {
		t.Errorf("Resolve(\"missing\") = (%q, %t), want (\"\", false)", got, ok)
	}
}

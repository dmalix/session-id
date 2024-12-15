package session_id

import (
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestGenerateSessionId(t *testing.T) {
	userId, err := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	hash, err := Generate(userId)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(hash) != 64 {
		t.Errorf("Expected hash length of 64, got: %d", len(hash))
	}
}

func TestIsValidTransformedHash(t *testing.T) {
	userId, err := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	hash, err := Generate(userId)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	log.Print(hash)

	err = Check(hash)
	if err != nil {
		t.Errorf("Expected hash to be valid")
	}

	invalidHash := "0123456789abcdef0123456789abcdef"
	err = Check(invalidHash)
	if err == nil {
		t.Errorf("Expected hash to be invalid")
	}
}

func BenchmarkGenerateSessionId(b *testing.B) {
	userId, err := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		b.Errorf("Unexpected error: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Generate(userId)
	}
}

func BenchmarkIsValidSessionId(b *testing.B) {
	userId, err := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		b.Errorf("Unexpected error: %v", err)
	}
	hash, _ := Generate(userId)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Check(hash)
	}
}

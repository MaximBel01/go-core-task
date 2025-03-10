package main

import (
	"testing"
)

func TestVariableToString(t *testing.T) {
	tests := []struct {
		input any
		want  string
	}{
		{42, "42"},
		{052, "42"},
		{0x2A, "42"},
		{6.9, "6.9"},
		{"Golang", "Golang"},
		{true, "true"},
		{complex(float32(1), float32(2)), "(1+2i)"},
	}

	for _, tt := range tests {
		got := variableToString(tt.input)
		if got != tt.want {
			t.Errorf("variableToString(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestConcatenateVariables(t *testing.T) {
	variables := []any{
		42,
		052,
		0x2A,
		3.14,
		"Golang",
		true,
		complex(float32(1), float32(2)),
	}

	want := "4242423.14Golangtrue(1+2i)"
	got := concatenateVariables(variables)

	if got != want {
		t.Errorf("concatenateVariables() = %v, want %v", got, want)
	}
}

func TestAddSaltAndHash(t *testing.T) {
	t.Run("Test with known input", func(t *testing.T) {
		runes := []rune("ab")
		salt := "go-2024"
		expectedHash := "660c6d86f7ab65536524d169964156f3cede5328a2086a0babcae3dd0e99fda0"
		hash := addSaltAndHash(runes, salt)
		if hash != expectedHash {
			t.Errorf("hash = %s, want %s", hash, expectedHash)
		}
	})

	t.Run("Test hash length", func(t *testing.T) {
		runes := []rune("test")
		salt := "go-2024"
		hash := addSaltAndHash(runes, salt)
		if len(hash) != 64 {
			t.Errorf("expected hash length 64, got %d", len(hash))
		}
	})
}

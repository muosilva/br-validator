package tests

import (
	"testing"

	validator "github.com/muosilva/br-validator/validators"
)

func TestIsValidCNPJ(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected bool
	}{
		{
			name:     "CNPJ válido",
			cnpj:     "11222333000181",
			expected: true,
		},
		{
			name:     "CNPJ inválido",
			cnpj:     "11222333000182",
			expected: false,
		},
		{
			name:     "CNPJ inválido - todos iguais",
			cnpj:     "11111111111111",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validator.IsValidCNPJ(tt.cnpj); got != tt.expected {
				t.Errorf("IsValidCNPJ() = %v, want %v", got, tt.expected)
			}
		})
	}
}

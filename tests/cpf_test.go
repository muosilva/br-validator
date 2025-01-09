package tests

import (
	"testing"

	validator "github.com/muosilva/br-validator/validators"
)

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected bool
	}{
		{
			name:     "CPF válido",
			cpf:      "52998224725",
			expected: true,
		},
		{
			name:     "CPF inválido",
			cpf:      "52998224726",
			expected: false,
		},
		{
			name:     "CPF inválido - todos iguais",
			cpf:      "11111111111",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validator.IsValidCPF(tt.cpf); got != tt.expected {
				t.Errorf("IsValidCPF() = %v, want %v", got, tt.expected)
			}
		})
	}
}

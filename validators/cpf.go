package validator

import (
	"regexp"
	"strings"
)

func IsValidCPF(cpf string) bool {
	cpf = RemoveNonDigits(cpf)

	if len(cpf) != 11 {
		return false
	}

	if strings.Repeat(string(cpf[0]), 11) == cpf {
		return false
	}

	return validateCheckDigits(cpf)
}

func RemoveNonDigits(input string) string {
	reg, _ := regexp.Compile(`\D`)
	return reg.ReplaceAllString(input, "")
}

func validateCheckDigits(cpf string) bool {
	if !calculateCheckDigit(cpf[:9], cpf[9]) {
		return false
	}

	return calculateCheckDigit(cpf[:10], cpf[10])
}

func calculateCheckDigit(cpf string, checkDigit byte) bool {
	sum := 0
	for i, char := range cpf {
		digit := int(char - '0')
		sum += digit * (len(cpf) + 1 - i)
	}

	remainder := (sum * 10) % 11
	if remainder == 10 {
		remainder = 0
	}

	return byte(remainder+'0') == checkDigit
}

package validator

import (
	"strconv"
	"strings"

	utils "github.com/muosilva/br-validator/pkg"
)

func IsValidCNPJ(cnpj string) bool {
	cnpj = utils.RemoveNonDigits(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	if strings.Repeat(string(cnpj[0]), 11) == cnpj {
		return false
	}

	return calculateCnpjCheckDigit(cnpj)

}

func calculateCnpjCheckDigit(cnpj string) bool {
	weight1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weight2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	calculateDigit := func(cnpj string, weights []int) int {
		sum := 0
		for i, weight := range weights {
			digit, _ := strconv.Atoi(string(cnpj[i]))
			sum += digit * weight
		}

		remainder := sum % 11
		if remainder < 2 {
			return 0
		}
		return 11 - remainder
	}

	firstDigit := calculateDigit(cnpj[:12], weight1)
	secondDigit := calculateDigit(cnpj[:13], weight2)

	return byte(firstDigit+'0') == cnpj[12] && byte(secondDigit+'0') == cnpj[13]
}

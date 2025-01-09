package pkg

import "regexp"

func RemoveNonDigits(input string) string {
	reg, _ := regexp.Compile(`\D`)
	return reg.ReplaceAllString(input, "")
}

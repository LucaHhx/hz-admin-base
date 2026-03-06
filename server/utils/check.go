package utils

import "regexp"

func IsEightDigits(s string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, s)
	return match
}

func IsTranCode(s string) bool {
	match, _ := regexp.MatchString(`^\d{6}$`, s)
	return match
}

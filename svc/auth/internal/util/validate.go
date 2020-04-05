package util

import "regexp"

func IsEmail(s string) bool {
	re, _ := regexp.Compile("[^@]+@.+\\..+")
	return re.Match([]byte(s))
}

func IsNumeric(s string) bool {
	re, _ := regexp.Compile("^([0-9]+)$")
	return re.Match([]byte(s))
}

package util

import "regexp"

func IsEmail(s string) bool {
	re, _ := regexp.Compile("[^@]+@.+\\..+")
	return re.Match([]byte(s))
}

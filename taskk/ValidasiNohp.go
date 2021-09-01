package main7

import (
	"regexp"
)

func ValidasiNohp(nohp string) string {
	number := regexp.MustCompile(`^(\+62|62)?[\s-]?0?8[1-9]{1}\d{1}[\s-]?\d{4}[\s-]?\d{2,5}$`)
	m := regexp.MustCompile("^(.*?)0(.*)$")
	str := "${1}62$2"
	res := m.ReplaceAllString(nohp, str)

	if number.MatchString(nohp) {
		return res
	}
	return "ERROR"
}

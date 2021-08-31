package main7

import (
	"github.com/dongri/phonenumber"
)

func ValidasiNohp(nohp string) string {
	number := phonenumber.Parse(nohp, "id")
	return number
}

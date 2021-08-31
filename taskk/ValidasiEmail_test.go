package main7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	bener    = true
	salah    = false
	pesanbnr = "Email anda benar"
	pesanslh = "Email anda salah"
)

func TestValidasiEmail(t *testing.T) {
	a := validateEmail("irfan@gmail.com")
	assert.Equal(t, a, bener, pesanbnr)

	b := validateEmail("irfan@ums.com")
	assert.Equal(t, b, bener, pesanbnr)

	c := validateEmail("irfan@gmail_com")
	assert.Equal(t, c, salah, pesanslh)

	d := validateEmail("irfan@gmailcom")
	assert.Equal(t, d, salah, pesanslh)

	e := validateEmail("irfangmail.com")
	assert.Equal(t, e, salah, pesanslh)

}

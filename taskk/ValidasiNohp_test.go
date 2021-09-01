package main7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	benerr    = "6281123321456"
	salahh    = "ERROR"
	pesanbnrr = "No hp yang di masukkan benar"
	pesanslhh = "No hp yang di masukkan salah"
)

func TestValidateNohp(t *testing.T) {
	a := ValidasiNohp("6281123321456")
	assert.Equal(t, a, benerr, pesanbnrr)

	b := ValidasiNohp("081123321456")
	assert.Equal(t, b, benerr, pesanbnrr)

	c := ValidasiNohp("081asdqwer")
	assert.Equal(t, c, salahh, pesanslhh)

	d := ValidasiNohp("081123456781234512321")
	assert.Equal(t, d, salahh, pesanslhh)

	e := ValidasiNohp("081213")
	assert.Equal(t, e, salahh, pesanslhh)

}

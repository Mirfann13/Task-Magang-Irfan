package main7

import (
	"fmt"
	"regexp"
)

func validateEmail(email string) bool {
	valEm := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return valEm.MatchString(email)
}

func main() {

	email := "abcd@gmail.com"

	if !validateEmail(email) {
		fmt.Println(email, "(INVALID EMAIL ADDRESS)")
	} else {
		fmt.Println(email, "(VALID EMAIL ADDRESS)")
	}

}

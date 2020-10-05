package utils

import (
	"math/rand"
	"regexp"
	"strings"
)

// ExtractNumbers extract numbers from string
func ExtractNumbers(str string) string {
	rx := regexp.MustCompile("[0-9]+")
	return strings.Join(rx.FindAllString(str, -1), "")
}

// RandomStringFromCharset generate random string from character list
func RandomStringFromCharset(n uint, letters string) (res string, err error) {
	bytes := make([]byte, n)
	_, err = rand.Read(bytes)
	if err != nil {
		return
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	res = string(bytes)
	return
}

// RandomString generate random string
func RandomString(n uint) (string, error) {
	return RandomStringFromCharset(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")
}

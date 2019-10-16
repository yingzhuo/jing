package jstr

import (
	"math/rand"
)

var (
	alphanumeric = ToRunes("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	alphabetic   = ToRunes("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numeric      = ToRunes("0123456789")
)

func NewRandomAlphabetic(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphabetic[rand.Intn(len(alphabetic))]
	}
	return string(b)
}

func NewRandomNumeric(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numeric[rand.Intn(len(numeric))]
	}
	return string(b)
}

func NewRandomAlphanumeric(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}
	return string(b)
}

func NewRandomString(n int, charset string) string {

	charsetLen := len(charset)
	if charsetLen == 0 || charsetLen > 256 {
		panic("charset n must be greater than 0 and less than or equal to 256")
	}
	var bitLength byte
	var bitMask byte
	for bits := charsetLen - 1; bits != 0; {
		bits = bits >> 1
		bitLength++
	}
	bitMask = 1<<bitLength - 1

	bufferSize := n + n/3

	result := make([]byte, n)
	for i, j, randomBytes := 0, 0, []byte{}; i < n; j++ {
		if j%bufferSize == 0 {
			// Random byte buffer is empty, get a new one
			randomBytes = secureRandomBytes(bufferSize)
		}
		// Mask bytes to get an index into the character slice
		if idx := int(randomBytes[j%n] & bitMask); idx < charsetLen {
			result[i] = charset[idx]
			i++
		}
	}

	return string(result)
}

func secureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, _ = rand.Read(randomBytes)
	return randomBytes
}

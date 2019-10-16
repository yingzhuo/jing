package jstr

import (
	"encoding/base64"
)

func Base64StdEncode(s string) string {
	input := []byte(s)
	return base64.StdEncoding.EncodeToString(input)
}

func Base64URLEncode(s string) string {
	input := []byte(s)
	return base64.URLEncoding.EncodeToString(input)
}

func Base64StdDecode(s string) string {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func Base64URLDecode(s string) string {
	data, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(data)
}

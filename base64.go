package main

import (
	"encoding/base64"
	"log"
)

// EncodeBase64
func EncodeBase64(message []byte) string {
	return base64.StdEncoding.EncodeToString(message)
}

// DecodeBase64
func DecodeBase64(message string) []byte {
	data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		log.Fatal("error when decoding base64 string '"+message+"': ", err)
	}
	return data
}

package main

import (
	"bytes"
	"errors"
)

// PKCS5Padding the byte array with blockSize
func PKCS5Padding(src []byte, blockSize int) []byte {
	srcLen := len(src)
	padLen := blockSize - (srcLen % blockSize)
	padText := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(src, padText...)
}

// PKCS5Unpadding the byte array
func PKCS5Unpadding(src []byte, blockSize int) ([]byte, error) {
	srcLen := len(src)
	paddingLen := int(src[srcLen-1])
	if paddingLen >= srcLen || paddingLen > blockSize {
		return nil, errors.New("byte array too short for unpadding")
	}
	return src[:srcLen-paddingLen], nil
}

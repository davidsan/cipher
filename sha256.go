package main

import (
	"crypto/sha256"
)

// ComputeSHA256 hash and truncate the bye array
func ComputeSHA256(msg []byte, len int) []byte {
	h := sha256.New()
	h.Write(msg)
	res := make([]byte, len)
	copy(res, h.Sum(nil))
	return res
}

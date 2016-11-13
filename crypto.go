package main

import (
	"crypto/cipher"
	"fmt"
	"log"
)

// Encrypt the message
func Encrypt(encryptor cipher.BlockMode, msg []byte) (cipherText []byte, err error) {
	if msg != nil {
		msg = PKCS5Padding(msg, encryptor.BlockSize())
		if len(msg) < encryptor.BlockSize() || len(msg)%encryptor.BlockSize() != 0 {
			log.Fatal("error:", err)
			return
		}
		cipherText = msg
		encryptor.CryptBlocks(cipherText, msg)
	}
	return
}

// Decrypt the ciphertext
func Decrypt(decryptor cipher.BlockMode, cipherText []byte) (msg []byte, err error) {
	if decryptor != nil {
		if len(cipherText) < decryptor.BlockSize() || len(cipherText)%decryptor.BlockSize() != 0 {
			fmt.Println("length error")
			return
		}
		msg = cipherText
		decryptor.CryptBlocks(msg, cipherText)
		msg, err = PKCS5Unpadding(msg, decryptor.BlockSize())
	}
	return
}

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

func main() {

	cfg := ParseConf()
	derivedKey := pbkdf2.Key([]byte(cfg.key), []byte(cfg.salt), cfg.iter, cfg.keyLen, sha256.New)
	iv := ComputeSHA256([]byte(cfg.vector), cfg.keyLen)

	var c cipher.Block
	var err error
	if c, err = aes.NewCipher(derivedKey); err != nil {
		log.Fatal("error on creating cipher:", err)
	}

	if cfg.decrypt {
		dbm := cipher.NewCBCDecrypter(c, iv)
		var decryptedData []byte
		if decryptedData, err = Decrypt(dbm, DecodeBase64(cfg.msg)); err != nil {
			log.Fatal("error on decrypt:", err)
		}
		fmt.Println(string(decryptedData))

	} else {
		ebm := cipher.NewCBCEncrypter(c, iv)
		var encryptedData []byte
		if encryptedData, err = Encrypt(ebm, []byte(cfg.msg)); err != nil {
			log.Fatal("error on encrypt:", err)
		}
		fmt.Println(EncodeBase64(encryptedData))
	}
}

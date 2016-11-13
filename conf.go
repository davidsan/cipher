package main

import "flag"
import "io/ioutil"
import "log"

// Configuration of the app
type Configuration struct {
	msg     string
	key     string
	vector  string
	salt    string
	keyLen  int
	iter    int
	decrypt bool
}

// ParseConf for the app
func ParseConf() Configuration {

	msgPtr := flag.String("msg", "false", "Message to encrypt / decrypt")
	keyFilePtr := flag.String("key", "key.txt", "Key file")
	vectorFilePtr := flag.String("iv", "iv.txt", "Initialization vector file")
	saltFilePtr := flag.String("salt", "salt.txt", "Salt file")

	keyLenPtr := flag.Int("keyLen", 16, "Key length to use (for PBKDF2 and AES)")
	iterPtr := flag.Int("iter", 100000, "Number of iterations (for PBKDF2)")
	decryptPtr := flag.Bool("decrypt", false, "If set will decrypt, else encrypt")

	flag.Parse()

	if len(*keyFilePtr) == 0 {
		log.Fatal("error: key file required")
	}
	if len(*vectorFilePtr) == 0 {
		log.Fatal("error: vector file required")
	}
	if len(*saltFilePtr) == 0 {
		log.Fatal("error: salt file required")
	}

	// attempt to read the File
	var buf []byte
	var err error

	if buf, err = ioutil.ReadFile(*keyFilePtr); err != nil {
		log.Fatal("error with key file:", err)
	}
	key := string(buf)

	if buf, err = ioutil.ReadFile(*vectorFilePtr); err != nil {
		log.Fatal("error with vector file:", err)
	}
	vector := string(buf)

	if buf, err = ioutil.ReadFile(*saltFilePtr); err != nil {
		log.Fatal("error with salt file:", err)
	}
	salt := string(buf)

	return Configuration{
		msg:     *msgPtr,
		key:     key,
		vector:  vector,
		salt:    salt,
		keyLen:  *keyLenPtr,
		iter:    *iterPtr,
		decrypt: *decryptPtr,
	}
}

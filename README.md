# cipher

## Install

```
$ go install github.com/davidsan/cipher
```

## Usage

```
$ $GOPATH/bin/cipher -msg "hello world"
8uGE5mxH+e0NYp6frQQjOQ==

$ $GOPATH/bin/cipher -msg "8uGE5mxH+e0NYp6frQQjOQ==" -decrypt
hello world

$ $GOPATH/bin/cipher -help
Usage of bin/cipher:
  -decrypt
    	If set will decrypt, else encrypt
  -iter int
    	Number of iterations (for PBKDF2) (default 100000)
  -iv string
    	Initialization vector file (default "iv.txt")
  -key string
    	Key file (default "key.txt")
  -keyLen int
    	Key length to use (for PBKDF2 and AES) (default 16)
  -msg string
    	Message to encrypt / decrypt
  -salt string
    	Salt file (default "salt.txt")
```
  

## Notes
Key, IV and salt can be generated with `openssl` for example

```
$ openssl rand -base64 1024 > key.txt
$ openssl rand -base64 1024 > iv.txt
$ openssl rand -base64 1024 > salt.txt
```

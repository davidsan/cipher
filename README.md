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
```
  

## Notes
Key, IV and salt can be generated with `openssl` for example

```
$ openssl rand -base64 1024 > key.txt
$ openssl rand -base64 1024 > iv.txt
$ openssl rand -base64 1024 > salt.txt
```

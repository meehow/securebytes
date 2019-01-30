# Secure Bytes [![GoDoc](https://godoc.org/github.com/github.com/meehow/securebytes?status.svg)](http://godoc.org/github.com/meehow/securebytes) [![Go Report Card](https://goreportcard.com/badge/github.com/meehow/securebytes)](https://goreportcard.com/report/github.com/meehow/securebytes#SecureBytes)

Secure Bytes takes any Go data type, serializes it to JSON or GOB and encrypts it.
The goal of this library is to generate smaller cookies than [securecookie](https://github.com/gorilla/securecookie) does.
It's achieved by using [Authenticated Encryption](https://en.wikipedia.org/wiki/Authenticated_encryption) instead of HMAC.

## Installation

```
go get -u github.com/meehow/securebytes
```

## Usage

Please check [example](examples/cookie.go) to check how to use `EncryptToBase64`
and `DecryptBase64`.

If you need plain `[]byte` output, you can use `Encrypt` and `Decrypt` functions instead.

You can find more informations in the [documentation](https://godoc.org/github.com/meehow/securebytes#SecureBytes).

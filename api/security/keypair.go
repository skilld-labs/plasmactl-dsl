package security

import "fmt"

type KeyPair struct {
	Encryption
}

type PrivateKeyEmptyError struct {
	msg string
}

func NewPrivateKeyEmptyError() PrivateKeyEmptyError {
	return PrivateKeyEmptyError{
		msg: "no private key provided",
	}
}

func (e PrivateKeyEmptyError) Error() string {
	return e.msg
}

type PrivateKeyInvalidFormat struct {
	msg string
}

func NewPrivateKeyInvalidFormatError(key string) PrivateKeyInvalidFormat {
	return PrivateKeyInvalidFormat{
		msg: fmt.Sprintf("private key provided has invalid format '%s'", key),
	}
}

func (e PrivateKeyInvalidFormat) Error() string {
	return e.msg
}

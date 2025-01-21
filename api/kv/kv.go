package kv

import (
	"dsl/api/security"
	"fmt"
)

type KV interface {
	Get(key string) (string, error)
	Set(key, value string) error
	GetBatch(keys []string) (KV, error)
	SetBatch(kv KV) error
	GetKeys() []string
}

type SealedKV interface {
	KV
	security.Encryption
}

// Defenition of kv errors
// KVKeyNotExistsError
type KVKeyNotExistsError struct {
	key string
}

func NewKVKeyNotExistsError(key string) KVKeyNotExistsError {
	return KVKeyNotExistsError{
		key: key,
	}
}

func (e KVKeyNotExistsError) Error() string {
	return fmt.Sprintf("kv key '%s' not exists", e.key)
}

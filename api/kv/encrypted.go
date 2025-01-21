package kv

import (
	"dsl/api/security"
	"errors"
)

type EncryptedKV struct {
	kv         KV
	encryption security.Encryption
}

func NewEncryptedKV(kv KV, encryption security.Encryption) *EncryptedKV {
	return &EncryptedKV{
		kv:         kv,
		encryption: encryption,
	}
}

func (v *EncryptedKV) Get(key string) (string, error) {
	encrypted, err := v.kv.Get(key)
	if errors.Is(err, KVKeyNotExistsError{}) {
		return "", NewKVKeyNotExistsError(key)
	} else if err != nil {
		return "", err
	}

	decrypted, err := v.Decrypt([]byte(encrypted))
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (v *EncryptedKV) Set(key, value string) error {
	var encrypted = []byte("")
	var err error

	if value != "" {
		// validation hook isAgeArmor
		// NewAlreadyEncryptedError
		if encrypted, err = v.Encrypt([]byte(value)); err != nil {
			return err
		}
	}

	if err = v.kv.Set(key, string(encrypted)); err != nil {
		return err
	}
	return nil
}

func (v *EncryptedKV) GetBatch(keys []string) (KV, error) {
	kv, err := v.kv.GetBatch(keys)
	if err != nil {
		return nil, err
	}
	return kv, nil
}

func (v *EncryptedKV) SetBatch(kv KV) error {
	if err := v.kv.SetBatch(kv); err != nil {
		return err
	}
	return nil
}

func (v *EncryptedKV) GetKeys() []string {
	return v.GetKeys()
}

func (v *EncryptedKV) Encrypt(data []byte) ([]byte, error) {
	encrypted, err := v.encryption.Encrypt(data)
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

func (v *EncryptedKV) Decrypt(data []byte) ([]byte, error) {
	decrypted, err := v.encryption.Decrypt(data)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

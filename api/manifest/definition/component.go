package definition

import (
	"bytes"
	"dsl/api/kv"
	"dsl/api/security"
	"fmt"
	"gopkg.in/yaml.v3"
)

type Manifest[T Schema] struct {
	Schema    T `yaml:",inline" validate:"required"`
	validator *Validator[T]
}

func NewManifest[T Schema]() (*Manifest[T], error) {
	validator, err := NewValidator[T]()
	manifest := &Manifest[T]{
		validator: validator,
	}
	if err != nil {
		return nil, fmt.Errorf("unable to create validator for type: '%+v'\ndue to error '%s'\n", manifest.Schema, err)
	}
	return manifest, nil
}

func (m *Manifest[T]) Unmarshal(deserialized []byte) error {
	reader := bytes.NewReader(deserialized)
	yamlDecoder := yaml.NewDecoder(reader)
	yamlDecoder.KnownFields(true)
	if err := yamlDecoder.Decode(&m.Schema); err != nil {
		return err
	}
	return nil
}

func (m *Manifest[T]) Marshal() ([]byte, error) {
	serialized, err := yaml.Marshal(&m.Schema)
	if err != nil {
		return nil, err
	}
	return serialized, err
}

func (m *Manifest[T]) Validate() error {
	if err := m.validator.Validate(&m.Schema); err != nil {
		return err
	}
	return nil
}

func (m *Manifest[T]) GetSecrets(encryption security.Encryption) kv.KV {
	secretsKV := m.Schema.GetCore().GetSecrets()
	secrets := kv.PlainKVFromMap(secretsKV)
	if encryption == nil {
		return secrets
	}
	return kv.NewEncryptedKV(secrets, encryption)
}

func (m *Manifest[T]) GetSecret(key string, encryption security.Encryption) (string, error) {
	secrets := m.GetSecrets(encryption)
	secret, err := secrets.Get(key)
	if err != nil {
		return "", err
	}
	return secret, nil
}

func (m *Manifest[T]) SetSecret(key, value string, encryption security.Encryption) error {
	secrets := m.GetSecrets(encryption)
	return secrets.Set(key, value)
}

func (m *Manifest[T]) UpdateSecrets(kv kv.KV, encryption security.Encryption) error {
	secrets := m.GetSecrets(encryption)
	return secrets.SetBatch(kv)
}

package definition

import (
	"dsl/api/kv"
	"dsl/api/schema/plasma/v1"
	"dsl/api/security"
)

type ComponentMarshaller interface {
	Marshal() ([]byte, error)
	Unmarshal(serialized []byte) error
}

type ComponentValidator interface {
	Validate() error
}

type ComponentSecretKV interface {
	GetSecrets(encryption security.Encryption) kv.KV
	GetSecret(key string, encryption security.Encryption) (string, error)
	SetSecret(key, value string, encryption security.Encryption) error
	UpdateSecrets(kv kv.KV, encryption security.Encryption) error
}

type Component interface {
	ComponentMarshaller
	ComponentValidator
	ComponentSecretKV
}

type HasCore interface {
	GetCore() *v1.Core
}

type Schema interface {
	v1.Application | v1.Entity | v1.Flow | v1.Function | v1.Library | v1.Metric | v1.Model | v1.Node | v1.Package | v1.Platform | v1.Service | v1.Skill | v1.Software
	HasCore
}

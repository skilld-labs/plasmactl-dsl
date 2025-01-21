package kv

import (
	"fmt"
)

type PlainKV struct {
	kv            map[string]string
	AllowOverride bool
}

type Opts func(kv *PlainKV)

func OptOverride(kv *PlainKV) {
	kv.AllowOverride = true
}

func DefaultOpts(kv *PlainKV) {
	kv.kv = make(map[string]string)
	kv.AllowOverride = false
}

func NewPlainKV(opts ...Opts) *PlainKV {
	var kv *PlainKV
	DefaultOpts(kv)
	for _, setOpt := range opts {
		setOpt(kv)
	}
	return kv
}

func PlainKVFromMap(fromMap map[string]string, opts ...Opts) KV {
	kv := &PlainKV{kv: fromMap}
	for _, setOpt := range opts {
		setOpt(kv)
	}
	return kv
}

func (p *PlainKV) Get(key string) (string, error) {
	kv, exists := p.kv[key]
	if !exists {
		return "", NewKVKeyNotExistsError(key)
	}
	return kv, nil
}

func (p *PlainKV) Set(key string, value string) error {
	if value == "" {
		delete(p.kv, key)
	}
	if valueKV, exists := p.kv[key]; exists && valueKV != value && !p.AllowOverride {
		return fmt.Errorf("key '%s' in KV with defferent value assignment found. Overriding prohibited by opt", key)
	}
	p.kv[key] = value
	return nil
}

func (p *PlainKV) GetBatch(keys []string) (KV, error) {
	var kv PlainKV
	for _, key := range keys {
		val, err := p.Get(key)
		if err != nil {
			return nil, err
		}
		if err := kv.Set(key, val); err != nil {
			return nil, err
		}
	}
	return &kv, nil
}

func (p *PlainKV) SetBatch(kv KV) error {
	for _, key := range kv.GetKeys() {
		val, err := kv.Get(key)
		if err != nil {
			return err
		}
		if err := p.Set(key, val); err != nil {
			return err
		}
	}

	return nil
}

func (p *PlainKV) GetKeys() []string {
	keys := make([]string, len(p.kv))
	i := 0
	for key := range p.kv {
		keys[i] = key
		i++
	}
	return keys
}

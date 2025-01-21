package service

import (
	"dsl/api/fs"
	"dsl/api/kv"
	"dsl/api/manifest/definition"
	"dsl/api/reader"
	"dsl/api/security"
	"errors"
	"fmt"
	"os"
)

type ManifestService struct {
	manifests map[string]definition.Component
}

func (s *ManifestService) ReadManifestFiles(paths []string, ext string) error {
	files, err := fs.DiscoverFiles(paths, ext)
	if err != nil {
		return err
	}

	s.manifests = make(map[string]definition.Component, len(files))

	for _, file := range files {
		manifest, err := s.readManifestFile(file)
		if err != nil {
			return err
		}

		s.manifests[file] = manifest
	}

	return nil
}

func (s *ManifestService) readManifestFile(file string) (definition.Component, error) {
	serialized, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error while reading '%s': '%s'", file, err.Error())
	}

	component, err := reader.ParseComponent(serialized)
	if err != nil {
		return nil, fmt.Errorf("error while creating component from file '%s':\n%s\n", file, err.Error())
	}

	if err := component.Unmarshal(serialized); err != nil {
		return nil, fmt.Errorf("error '%s' while deserializing '%s' with following content:\n```\n%s\n```", err.Error(), file, string(serialized))
	}

	if err := component.Validate(); err != nil {
		return nil, fmt.Errorf("validation error invalid component file '%s': '%s'", file, err.Error())
	}

	return component, nil
}

func (s *ManifestService) SaveManifestFile(file string) error {
	manifest, ok := s.manifests[file]
	if !ok {
		return fmt.Errorf("manifest save error. No such manifest file '%s'", file)
	}

	if err := manifest.Validate(); err != nil {
		return fmt.Errorf("validation error invalid component file '%s': '%s'", file, err.Error())
	}

	serialized, err := manifest.Marshal()
	if err != nil {
		return fmt.Errorf("error '%s' while serializing '%s' with following content:\n```\n%s\n```", err.Error(), file, string(serialized))
	}

	if err := fs.WriteFile(file, serialized); err != nil {
		return fmt.Errorf("error while saving manifest. Error writing to file '%s' of content:\n%s\nerror processing data: '%s'", file, err, err.Error())
	}

	return nil
}

func (s *ManifestService) GetManifests() map[string]definition.Component {
	return s.manifests
}

func (s *ManifestService) GetSecret(key string, encryption security.Encryption) (string, error) {
	for _, manifest := range s.manifests {
		secrets := manifest.GetSecrets(encryption)

		secret, err := secrets.Get(key)
		if !errors.Is(err, kv.KVKeyNotExistsError{}) && err != nil {
			return "", err
		} else {
			return secret, nil
		}
	}

	return "", kv.NewKVKeyNotExistsError(key)
}

func (s *ManifestService) SetSecret(key, value string, encryption security.Encryption, save bool) error {
	for file, manifest := range s.manifests {
		secrets := manifest.GetSecrets(encryption)

		if err := secrets.Set(key, value); err != nil {
			return err
		}

		if err := manifest.UpdateSecrets(secrets, encryption); err != nil {
			return err
		}

		if save {
			if err := s.SaveManifestFile(file); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *ManifestService) GetSecretList(keys []string, encryption security.Encryption) (kv.KV, error) {
	var secretKV kv.KV
	for _, manifest := range s.manifests {
		secrets := manifest.GetSecrets(encryption)

		batch, err := secrets.GetBatch(keys)
		if !errors.Is(err, kv.KVKeyNotExistsError{}) && err != nil {
			return nil, err
		}

		if err := secretKV.SetBatch(batch); err != nil {
			return nil, err
		}
	}

	return secretKV, nil
}

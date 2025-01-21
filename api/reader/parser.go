package reader

import (
	"dsl/api/manifest/definition"
	v1 "dsl/api/schema/plasma/v1"
	"fmt"
	"gopkg.in/yaml.v3"
)

func ParseComponent(serialized []byte) (definition.Component, error) {
	kind, err := parseKind(serialized)
	if err != nil {
		return nil, err
	}
	switch kind {
	case definition.KindApplication:
		return definition.NewManifest[v1.Application]()
	case definition.KindEntity:
		return definition.NewManifest[v1.Entity]()
	case definition.KindFlow:
		return definition.NewManifest[v1.Flow]()
	case definition.KindFunction:
		return definition.NewManifest[v1.Function]()
	case definition.KindLibrary:
		return definition.NewManifest[v1.Library]()
	case definition.KindMetric:
		return definition.NewManifest[v1.Metric]()
	case definition.KindModel:
		return definition.NewManifest[v1.Model]()
	case definition.KindNode:
		return definition.NewManifest[v1.Node]()
	case definition.KindPackage:
		return definition.NewManifest[v1.Package]()
	case definition.KindPlatform:
		return definition.NewManifest[v1.Platform]()
	case definition.KindService:
		return definition.NewManifest[v1.Service]()
	case definition.KindSkill:
		return definition.NewManifest[v1.Skill]()
	case definition.KindSoftware:
		return definition.NewManifest[v1.Software]()
	default:
		return nil, fmt.Errorf("not allowed component kind '%s' - incorrect manifest definition:\n%s\n", kind, string(serialized))
	}
}

func parseKind(serialized []byte) (string, error) {
	kindField := &struct {
		Kind string `yaml:"kind" validate:"isKind"`
	}{
		Kind: "",
	}

	if err := yaml.Unmarshal(serialized, kindField); err != nil {
		return "", fmt.Errorf("unable to parse kind from content:\n%s\n", string(serialized))
	}

	if kindField.Kind == "" {
		return "", fmt.Errorf("parsed empty kind from content:\n%s\n", string(serialized))
	}

	return kindField.Kind, nil
}

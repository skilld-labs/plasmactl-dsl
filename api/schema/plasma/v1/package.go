package v1

type Package struct {
	Core            `yaml:",inline" validate:"required"`
	PackageSettings `yaml:"settings" validate:"required"`
}

type PackageSettings struct {
	PackageType string `yaml:"type"   validate:"required"`
	PackageTag  string `yaml:"tag"    validate:"required"`
	PackageUrl  string `yaml:"url"    validate:"required"`
}

func (p Package) GetCore() *Core {
	return &p.Core
}

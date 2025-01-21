package v1

type Entity struct {
	Core           `yaml:",inline" validate:"required"`
	EntitySettings `yaml:"settings" validate:"required"`
}

type EntitySettings struct {
	SchemaVersion string `yaml:"schemaVersion"`
}

func (e Entity) GetCore() *Core {
	return &e.Core
}

package v1

type Function struct {
	Core             `yaml:",inline" validate:"required"`
	FunctionSettings `yaml:"settings" validate:"required"`
}

type FunctionSettings struct {
	OperationMode string `yaml:"operationMode" validate:"isConcurrencyMode"`
}

func (f Function) GetCore() *Core {
	return &f.Core
}

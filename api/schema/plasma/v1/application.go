package v1

type Application struct {
	Core                `yaml:",inline" validate:"required"`
	ApplicationSettings `yaml:"settings" validate:"required"`
}

type ApplicationSettings struct {
	MaxRetries uint `yaml:"maxRetries"`
}

func (a Application) GetCore() *Core {
	return &a.Core
}

package v1

type Software struct {
	Core             `yaml:",inline" validate:"required"`
	SoftwareSettings `yaml:"settings" validate:"required"`
}

type SoftwareSettings struct {
	ServiceBuilderContext string `yaml:"serviceBuilderContextType" validate:"isServiceBuilderCtx"`
	ServiceScope          string `yaml:"serviceScope"              validate:"isServiceScope"`
	Port                  uint16 `yaml:"port"                      validate:"isTCPv4Port"`
}

func (s Software) GetCore() *Core {
	return &s.Core
}

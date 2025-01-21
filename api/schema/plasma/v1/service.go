package v1

type Service struct {
	Core            `yaml:",inline" validate:"required"`
	ServiceSettings `yaml:"settings" validate:"required"`
}

type ServiceSettings struct {
	ServiceBuilderContext string `yaml:"serviceBuilderContextType"  validate:"isServiceBuilderCtx"`
	ServiceScope          string `yaml:"serviceScope"               validate:"isServiceScope"`
	RetryLimit            uint   `yaml:"retryLimit"`
	GrafanaApiKey         string `yaml:"grafanaApiKey"`
}

func (s Service) GetCore() *Core {
	return &s.Core
}

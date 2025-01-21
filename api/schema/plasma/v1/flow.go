package v1

type Flow struct {
	Core           `yaml:",inline" validate:"required"`
	FlowAttributes `yaml:",inline" validate:"required"`
	FlowSettings   `yaml:"settings" validate:"required"`
}

type FlowAttributes struct {
	FlowType   string `yaml:"flowType"      validate:"required,isFlowType"`
	FlowName   string `yaml:"flowName"      validate:"required"`
	FlowInput  string `yaml:"flowInput"     validate:"required"`
	FlowOutput string `yaml:"flowOutput"`
}

type FlowSettings struct {
	PostgresVersion string `yaml:"postgresVersion" validate:"semver"`
	GrafanaAlerting bool   `yaml:"grafanaAlerting"`
	Timeout         string `yaml:"timeout,omitempty" validate:"timeInterval"`
}

func (f Flow) GetCore() *Core {
	return &f.Core
}

package v1

type Metric struct {
	Core           `yaml:",inline" validate:"required"`
	MetricSettings `yaml:"settings" validate:"required"`
}

type MetricSettings struct {
	AggregationInterval string `yaml:"aggregationInterval" validate:"timeInterval"`
}

func (m Metric) GetCore() *Core {
	return &m.Core
}

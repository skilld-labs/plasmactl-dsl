package v1

type Platform struct {
	Core `yaml:",inline" validate:"required"`
}

func (p Platform) GetCore() *Core {
	return &p.Core
}

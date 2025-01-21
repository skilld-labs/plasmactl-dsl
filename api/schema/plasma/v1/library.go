package v1

type Library struct {
	Core            `yaml:",inline" validate:"required"`
	LibrarySettings `yaml:"settings" validate:"required"`
}

type LibrarySettings struct {
	PluginPath string `yaml:"pluginPath"`
}

func (l Library) GetCore() *Core {
	return &l.Core
}

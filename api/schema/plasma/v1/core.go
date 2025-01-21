package v1

type Core struct {
	Kind         string            `yaml:"kind"          validate:"isKind,eqcsfield=Metadata.Pck"`
	Metadata     Metadata          `yaml:"metadata"       validate:"required"`
	Name         string            `yaml:"name"          validate:"required,max=63,eqcsfield=Metadata.Pcn"`
	ShortName    string            `yaml:"shortName"     validate:"required,max=63,eqcsfield=Metadata.Pcsn"`
	Namespace    string            `yaml:"namespace"     validate:"required,max=63"`
	Owner        string            `yaml:"owner"         validate:"required"`
	APIVersion   string            `yaml:"apiVersion"    validate:"required"`
	Description  string            `yaml:"description"   validate:"required"`
	Dependencies []string          `yaml:"dependencies"`
	Hooks        Hooks             `yaml:"hooks"`
	Secrets      map[string]string `yaml:"secrets,flow"`
}

func (c *Core) GetSecrets() map[string]string {
	return c.Secrets
}

type Metadata struct {
	Pcn  string   `yaml:"pcn"             validate:"required,max=63"`
	Pcsn string   `yaml:"pcsn"            validate:"required,max=63"`
	Pck  string   `yaml:"pck"             validate:"required"`
	Pcl  string   `yaml:"pcl"             validate:"isPcl"`
	Pcc  string   `yaml:"pcc"             validate:"required,max=63"`
	Pcln string   `yaml:"pcln"            validate:"required,max=63"`
	Pct  []string `yaml:"pct,flow"`
	Pcv  string   `yaml:"pcv"             validate:"required,semver"`
	Pci  string   `yaml:"pci"`
}

type Hooks struct {
	Initialize          string `yaml:"initialize"`
	ResolveDependencies string `yaml:"resolveDependencies"`
	Build               string `yaml:"build"`
	AwaitReadiness      string `yaml:"awaitReadiness"`
	Configure           string `yaml:"configure"`
	Expose              string `yaml:"expose"`
	Decommission        string `yaml:"decommission"`
}

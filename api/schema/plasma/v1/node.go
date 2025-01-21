package v1

type Node struct {
	Core         `yaml:",inline" validate:"required"`
	NodeSettings `yaml:"settings" validate:"required"`
}

// NodeSettings holds settings specific to nodes
type NodeSettings struct {
	Region    string   `yaml:"region"`
	CPU       string   `yaml:"cpu"       validate:"milliCPU"`
	RAM       string   `yaml:"ram"       validate:"infoUnits"`
	Storage   string   `yaml:"storage"   validate:"infoUnits"`
	PrivateIP string   `yaml:"privateIP" validate:"ipv4|ipv6"`
	PublicIP  string   `yaml:"publicIP"  validate:"ipv4|ipv6"`
	Chassis   []string `yaml:"chassis"`
}

func (n Node) GetCore() *Core {
	return &n.Core
}

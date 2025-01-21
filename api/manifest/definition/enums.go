package definition

type Enum map[string]bool

func (e Enum) Exists(elem string) bool {
	if exists, ok := e[elem]; !exists || !ok {
		return false
	}
	return true
}

const (
	KindApplication = "Application"
	KindEntity      = "Entity"
	KindFlow        = "Flow"
	KindFunction    = "Function"
	KindLibrary     = "Library"
	KindMetric      = "Metric"
	KindModel       = "Model"
	KindNode        = "Node"
	KindPackage     = "Package"
	KindPlatform    = "Platform"
	KindService     = "Service"
	KindSkill       = "Skill"
	KindSoftware    = "Software"
)

var (
	EnumKind = Enum{
		KindApplication: true,
		KindEntity:      true,
		KindFlow:        true,
		KindFunction:    true,
		KindLibrary:     true,
		KindMetric:      true,
		KindModel:       true,
		KindNode:        true,
		KindPackage:     true,
		KindPlatform:    true,
		KindService:     true,
		KindSkill:       true,
		KindSoftware:    true,
	}

	EnumPcl = Enum{
		"Platform":      true,
		"Foundation":    true,
		"Integration":   true,
		"Interaction":   true,
		"Cognition":     true,
		"Conversation":  true,
		"Stabilization": true,
	}

	EnumFlowType = Enum{
		"lake":     true,
		"data":     true,
		"http":     true,
		"schedule": true,
		"event":    true,
	}

	EnumConcurrencyMode = Enum{
		"sync":  true,
		"async": true,
	}

	EnumSkillStage = Enum{
		"data":        true,
		"information": true,
		"knowledge":   true,
		"wisdom":      true,
	}

	EnumServiceBuilderContext = Enum{
		"cluster": true,
		"image":   true,
	}

	EnumServiceScope = Enum{
		"local":   true,
		"cluster": true,
	}
)

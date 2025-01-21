package v1

type Model struct {
	Core     `yaml:",inline"     validate:"required"`
	Settings ModelSettings      `yaml:"settings"`
	Platform PlatformComponents `yaml:"platform"`
	Packages []PackageRef       `yaml:"packages"`
}

func (m Model) GetCore() *Core {
	return &m.Core
}

type ModelSettings struct {
	GlobalMonitoringEnabled bool   `yaml:"globalMonitoringEnabled" validate:"required"`
	DefaultTheme            string `yaml:"defaultTheme"            validate:"required"`
}

type PlatformComponents struct {
	Foundation   FoundationComponent   `yaml:"foundation"`
	Interaction  InteractionComponent  `yaml:"interaction"`
	Integration  IntegrationComponent  `yaml:"integration"`
	Cognition    CognitionComponent    `yaml:"cognition"`
	Conversation ConversationComponent `yaml:"conversation"`
}

type FoundationComponent struct {
	Cluster       []ClusterComponent `yaml:"cluster"       validate:"required"`
	Network       NetworkComponent   `yaml:"network"`
	Storage       []string           `yaml:"storage"`
	Observability []string           `yaml:"observability" validate:"required"`
	Security      []string           `yaml:"security"`
	Core          []string           `yaml:"core"`
}

type ClusterComponent struct {
	Component       string          `yaml:"component" validate:"required"`
	ClusterSettings ClusterSettings `yaml:"clusterSettings"`
}

type ClusterSettings struct {
	NodeKind string `yaml:"nodeKind" validate:"required"`
}

type NetworkComponent struct {
	Ingress []string `yaml:"ingress"`
}

type InteractionComponent struct {
	Observability []ObservabilityComponent `yaml:"observability"`
	Management    []ManagementComponent    `yaml:"management"`
	Communication []string                 `yaml:"communication"`
	Operations    []string                 `yaml:"operations"`
	Marketing     []MarketingComponent     `yaml:"marketing"`
}

type ObservabilityComponent struct {
	Component string                `yaml:"component"`
	Settings  ObservabilitySettings `yaml:"settings" validate:"required"`
}

type ObservabilitySettings struct {
	GrafanaURL               string `yaml:"grafanaURL" validate:"required"`
	DashboardRefreshInterval string `yaml:"dashboardRefreshInterval" validate:"required"`
	MonitoringEnabled        bool   `yaml:"monitoringEnabled" validate:"required"`
	ActiveTheme              string `yaml:"activeTheme" validate:"required"`
}

type ManagementComponent struct {
	Component string             `yaml:"component" validate:"required"`
	Settings  ManagementSettings `yaml:"settings"`
}

type ManagementSettings struct {
	ERPDatabaseURL string `yaml:"erpDatabaseURL"`
}

type MarketingComponent struct {
	Component string            `yaml:"component"`
	Settings  MarketingSettings `yaml:"settings"`
}

type MarketingSettings struct {
	LandingPageTitle string `yaml:"landingPageTitle" validate:"required"`
	TrackVisitor     bool   `yaml:"trackVisitor"`
}

type IntegrationComponent struct {
	Core         []string `yaml:"core"`
	Organization []string `yaml:"organization"`
	Operations   []string `yaml:"operations"`
}

type CognitionComponent struct {
	Core      []string `yaml:"core"`
	Data      []string `yaml:"data"`
	Knowledge []string `yaml:"knowledge"`
	Wisdom    []string `yaml:"wisdom"`
}

type ConversationComponent struct {
	Core []string `yaml:"core"`
}

type PackageRef struct {
	Name   string          `yaml:"name"   validate:"required"`
	Source PackageSettings `yaml:"source" validate:"required"`
}

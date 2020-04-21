package core

//////// IAM OIDC CONFIGURATION

// IAMConfig : interface for generic IAM manager
type IAMConfig interface {
	Init() error
	RefreshAccessToken() (string, error)
	GetAccessToken() (string, error)
}

//////// INFRASTRUCTURE CONFIGURATION

// InfrastructureConfig the configuration for an infrastructure creation
type InfrastructureConfig struct {
	Nodes []NodeConfig `yaml:"nodes"`
}

// NodeConfig :
type NodeConfig struct {
	Image   string      `yaml:"host"`
	CPUs    string      `yaml:"CPUs,omitempty"`
	RAM     string      `yaml:"RAM,omitempty"`
	Flavor  string      `yaml:"flavor,omitempty"`
	Network NodeNetwork `yaml:"network"`
	Roles   []NodeRole  `yaml:"roles"`
	Labels  []NodeLabel `yaml:"labels,omitempty"`
	Count   int         `yaml:"count,omitempty"`
}

// NodeNetwork : describe network configuration for the node
type NodeNetwork struct {
	Ports       []int  `yaml:"ports"`
	IsPublic    bool   `yaml:"is_public"`
	NetworkID   string `yaml:"network_id,omitempty"`
	NetworkName string `yaml:"network_name,omitempty"`
}

// NodeLabel : label description of a node
type NodeLabel struct {
	Key   string
	Value string
}

// NodeRole : role description of a node
type NodeRole struct {
	string
}

// InfrastructureState descriptor
type InfrastructureState struct {
	labels []NodeLabel
}

//////// K8S CLUSTER CONFIGURATION

// ClusterConfig the configuration structure for k8s clusters
type ClusterConfig struct {
}

// LoadClusterConfig : load configuration file for K8s cluster
func LoadClusterConfig(configPath string) (ClusterConfig, error) {

	return ClusterConfig{}, nil
}

// ClusterState the k8s cluster state
type ClusterState struct {
}

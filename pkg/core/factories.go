package core

// InfrastructureManager is the interface for any plugin for cluster creation
type InfrastructureManager interface {
	// Create :
	Create() error
	// Update :
	Update() error
	// Scale :
	Scale(nodeLabel NodeLabel, num int) error
	// Destroy :
	Destroy() error
	// Status:
	Status() error
	// Describe
	Describe() error
}

// InfrastructureManagerFactory :
type InfrastructureManagerFactory interface {
	// Init : initialize factory
	Init() (*InfrastructureManager, error)
}

// ClusterManager is the interface for managing K8s clusters
type ClusterManager interface {
	// Create :
	Create() (ClusterState, error)
	// Update :
	Update() error
	// Scale :
	Scale(num int)
	// Destroy :
	Destroy() error
	// Status:
	Status() (string, error)
	// Describe
	Describe() (string, error)
}

// ClusterManagerFactory : factory method for the Cluster Manager
type ClusterManagerFactory interface {
	// Init : initialize a ClusterManager interface
	Init() (*ClusterManager, error)
}

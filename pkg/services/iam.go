package services

import "fmt"

// IndigoIAMConfig : configuration for Indigo-IAM
type IndigoIAMConfig struct {
	CredPath     string
	Name         string
	AccessToken  string
	ClientSecret string
	ClientID     string
	Endpoint     string
}

// Init initialize Indigo IAM configuration. As per core.IAMConfig interface
func (IndigoIAMConfig) Init() error {
	return fmt.Errorf("Operation not implemented")
}

// RefreshAccessToken initialize Indigo IAM configuration. As per core.IAMConfig interface
func (IndigoIAMConfig) RefreshAccessToken() (string, error) {
	return "", fmt.Errorf("Operation not implemented")
}

// GetAccessToken initialize Indigo IAM configuration. As per core.IAMConfig interface
func (IndigoIAMConfig) GetAccessToken() (string, error) {
	return "", fmt.Errorf("Operation not implemented")
}

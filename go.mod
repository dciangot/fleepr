module github.com/dciangot/fleepr

go 1.14

require (
	github.com/hashicorp/terraform v0.12.24
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rancher/rke v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.0.1+incompatible
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.2
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.6.5
)

module github.com/dciangot/fleepr

go 1.14

require (
	github.com/hashicorp/terraform v0.12.22
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rancher/rke v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
)

replace (
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.7
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.0.1+incompatible
	github.com/hashicorp/consul/api => github.com/hashicorp/consul/api v1.3.0
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.6.2-0.20191112221531-8742361660b6
)

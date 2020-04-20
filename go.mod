module github.com/dciangot/fleepr

go 1.14

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/rancher/rke v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	golang.org/x/crypto v0.0.0-20191202143827-86a70503ff7e // indirect
	golang.org/x/net v0.0.0-20191126235420-ef20fe5d7933 // indirect
	golang.org/x/tools v0.0.0-20191203134012-c197fd4bf371 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.0.1+incompatible
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.2
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.6.5
	github.com/terraform-providers/terraform-provider-openstack => github.com/terraform-providers/terraform-provider-openstack v1.27.0
)

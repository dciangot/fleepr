package services

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/dciangot/fleepr/pkg/core"
)

// ConfIM : configuration for Indigo InfrastructureManager auth
type ConfIM struct {
	ID       string `yaml:"id"`
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Token    string `yaml:"token,omitempty"`
}

// ConfCloud ...
type ConfCloud struct {
	ID            string `yaml:"id"`
	Type          string `yaml:"type"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	Host          string `yaml:"host"`
	Tenant        string `yaml:"tenant"`
	AuthURL       string `yaml:"auth_url,omitempty"`
	AuthVersion   string `yaml:"auth_version"`
	Domain        string `yaml:"domain,omitempty"`
	ServiceRegion string `yaml:"service_region,omitempty"`
}

// IMClientFactory : Indigo-dc InfrastructureManager client factory
type IMClientFactory struct {
	configPath string
	statePath  string
}

// IMClient : Indigo-dc InfrastructureManager client
type IMClient struct {
	ConfigIM    ConfIM
	ConfigCloud ConfCloud
	Context     core.InfrastructureConfig
	State       core.InfrastructureState
}

// Init : initialize IMClient
func (t IMClientFactory) Init() (*IMClient, error) {
	context := core.InfrastructureConfig{}

	// TODO: implement methods for IAM management
	configIM := ConfIM{}
	configCloud := ConfCloud{}

	return &IMClient{
		ConfigIM:    configIM,
		ConfigCloud: configCloud,
		Context:     context,
		State:       core.InfrastructureState{},
	}, nil
}

// Create : create infrastructure with IMClient
func (t IMClient) Create() error {
	tmpl := template.Must(template.New("IMClusterTemplate").Funcs(sprig.FuncMap()).Parse(IMClusterTemplate))

	var tmplBuffer bytes.Buffer
	err := tmpl.Execute(&tmplBuffer, t.Context)
	if err != nil {
		return fmt.Errorf("Failed to compile the template: %s", err)
	}

	headerCloudLine := map[string]string{}
	authHeaderCloud := core.PrepareAuthHeaders(headerCloudLine)
	fmt.Printf("HeaderCloud: %s", authHeaderCloud)

	headerIMLine := map[string]string{}
	authHeaderIM := core.PrepareAuthHeaders(headerIMLine)
	fmt.Printf("HeaderIM: %s", authHeaderIM)

	authHeaderList := []string{authHeaderCloud, authHeaderIM}
	authHeader := strings.Join(authHeaderList, "\\n")

	request := core.Request{
		URL:         t.ConfigIM.Host,
		RequestType: "POST",
		Headers: map[string]string{
			"Authorization": authHeader,
			"Content-Type":  "text/yaml",
		},
		Content: tmplBuffer.Bytes(),
	}

	body, statusCode, err := core.MakeRequest(request)
	if err != nil {
		return err
	}

	if statusCode == 200 {
		stringSplit := strings.Split(string(body), "/")
		fmt.Println("InfrastructureID: ", stringSplit[len(stringSplit)-1])
	} else {
		return fmt.Errorf("Error code %d: %s", statusCode, body)
	}

	_ = strings.Split(string(body), "/")
	// TODO: create .dodas dir and save infID

	return nil
}

// Update : update the deployment with the current Infrastructure configuration
func (t IMClient) Update() error {
	return fmt.Errorf("Operation not implemented")
}

// Scale : scale up or down selected node
func (t IMClient) Scale(nodeLabel core.NodeLabel, num int) error {
	return fmt.Errorf("Operation not implemented")
}

// Destroy : destroy the infrastructure
func (t IMClient) Destroy() error {
	return fmt.Errorf("Operation not implemented")
}

// Status : get status of the infrastructure
func (t IMClient) Status() error {
	return fmt.Errorf("Operation not implemented")
}

// Describe : get detailed information about the infrastructure
func (t IMClient) Describe() error {
	return fmt.Errorf("Operation not implemented")
}

// IMClusterTemplate : template TOSCA for IM
const IMClusterTemplate = `
tosca_definitions_version: tosca_simple_yaml_1_0

imports:
  - dodas_cod_types: https://raw.githubusercontent.com/dodas-ts/dodas-templates/master/tosca-types/dodas_types.yml

description: TOSCA template for a complete CMS computing cluster on top of K8s orchestrator

topology_template:

  inputs:

    number_of_masters:
      type: integer
      default: 1

    num_cpus_master: 
      type: integer
      default: 2

    mem_size_master:
      type: string
      default: "4 GB"

    number_of_slaves:
      type: integer
      default: {{ .Slaves.SlaveNum }}

    num_cpus_slave: 
      type: integer
      default: 4

    mem_size_slave:
      type: string
      default: "8 GB"

    server_image:
      type: string
      #default: "ost://openstack.fisica.unipg.it/cb87a2ac-5469-4bd5-9cce-9682c798b4e4"
      default: "ost://horizon.cloud.cnaf.infn.it/3d993ab8-5d7b-4362-8fd6-af1391edca39"
      #default: "ost://cloud.recas.ba.infn.it/1113d7e8-fc5d-43b9-8d26-61906d89d479"

    helm_cod_values: 
      type: string
      default: |
        externalIp:
          enabled: true
          ips:
            - {{ .ExternalIP }}
        gsi:
          enabled: {{ .GsiEnabled }}
          {{- if .GsiEnabled }}
          vo: {{ .VO }}
          vomses:
          - filename: {{ .VomsFile.Name }}
            content: | {{ .VomsFile.Content | indent 14 }}
          caCert:
            cert: | {{ .CacheCert | indent 14 }}
            key: | {{ .CacheKey | indent 14 }}
          proxy: true
          {{- end }}

        cache:
          {{- if .RedirectorHost }}
          redirHost: {{ .RedirectorHost }}
          {{- end }}
          originHost: {{ .OriginHost }}
          originXrdPort: {{ .OriginPort }}
        redirector:
          {{- if .RedirectorHost }}
          enabled: false
          {{- end}}
          service:
            cms:
              port: {{ .RedirectorPort }}
        proxy:
          {{- if .RedirectorHost }}
          enabled: false
          {{- end}}

  node_templates:

    helm_cod:
      type: tosca.nodes.DODAS.HelmInstall.CachingOnDemand
      properties:
        repos:
          - { name: dodas, url: "https://dodas-ts.github.io/helm_charts" }
        name: "cachingondemand"
        chart: "dodas/cachingondemand"
        kubeconfig_path: /var/lib/rancher/k3s/server/cred/admin.kubeconfig
        values_file: { get_input: helm_cod_values }
        externalIp: { get_attribute: [ k3s_master_server , public_address, 0 ]  }
      requirements:
        - host: k3s_master_server
        - dependency: k3s_slave_server

    k3s_master:
      type: tosca.nodes.DODAS.k3s 
      properties:
        master_ip: { get_attribute: [ k3s_master_server, private_address, 0 ] }
        mode: server
      requirements:
        - host: k3s_master_server

    k3s_slave:
      type: tosca.nodes.DODAS.k3s 
      properties:
        master_ip: { get_attribute: [ k3s_master_server, private_address, 0 ] }
        mode: node 
      requirements:
        - host: k3s_slave_server
        - dependency: k3s_master

    k3s_master_server:
      type: tosca.nodes.indigo.Compute
      capabilities:
        endpoint:
          properties:
            # network_name: infn-farm.PUBLIC
            network_name: PUBLIC
            ports:
              kube_port:
                protocol: tcp
                source: 30443
              xrd_port:
                protocol: tcp
                source: 31394
        scalable:
          properties:
            count: { get_input: number_of_masters }
        host:
          properties:
            # instance_type:  m1.medium
            num_cpus: { get_input: num_cpus_master }
            mem_size: { get_input: mem_size_master } 
        os:
          properties:
            image: { get_input: server_image }

    k3s_slave_server:
      type: tosca.nodes.indigo.Compute
      capabilities:
        endpoint:
          properties:
            #network_name: test-net.PRIVATE
            network_name: PRIVATE
        scalable:
          properties:
            count: { get_input: number_of_slaves }
        host:
          properties:
            #instance_type:  m1.medium
            num_cpus: { get_input: num_cpus_slave }
            mem_size: { get_input: mem_size_slave } 
        os:
          properties:
            # image: "ost://openstack.fisica.unipg.it/d9a41aed-3ebf-42f9-992e-ef0078d3de95"
            image: { get_input: server_image }

  outputs:
    k8s_endpoint:
      value: { concat: [ 'https://', get_attribute: [ k3s_master_server, public_address, 0 ], ':30443' ] }
`

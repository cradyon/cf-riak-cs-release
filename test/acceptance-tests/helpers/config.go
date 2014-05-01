package helpers

import (
	"encoding/json"
	"os"
)

type IntegrationConfig struct {
	AppsDomain        string `json:"apps_domain"`
	SystemDomain      string `json:"system_domain"`
	ApiEndpoint       string `json:"api"`

	RiakCsScheme      string `json:"riak_cs_scheme"`

	AdminUser         string `json:"admin_user"`
	AdminPassword     string `json:"admin_password"`

	ServiceName				string `json:"service_name"`
	PlanName					string `json:"plan_name"`

	SkipSSLValidation bool `json:"skip_ssl_validation"`
}

func LoadConfig() (config IntegrationConfig) {
	path := os.Getenv("CONFIG")
	if path == "" {
		panic("Must set $CONFIG to point to an integration config .json file.")
	}

	return LoadPath(path)
}

func LoadPath(path string) (config IntegrationConfig) {
	config = IntegrationConfig{
		SkipSSLValidation: false,
	}

	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	if config.ApiEndpoint == "" {
		panic("missing configuration 'api'")
	}

	if config.AdminUser == "" {
		panic("missing configuration 'admin_user'")
	}

	if config.ApiEndpoint == "" {
		panic("missing configuration 'admin_password'")
	}

	if config.ServiceName == "" {
		panic("missing configuration 'service_name'")
	}

	if config.PlanName == "" {
		panic("missing configuration 'plan_name'")
	}

	return
}

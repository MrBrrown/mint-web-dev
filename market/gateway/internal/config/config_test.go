package config_test

import (
	"fmt"
	"market/common/yamlconf"
	"marketapi/gateway/internal/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg := &config.Config{}
	err := yamlconf.Load("./test_config.yaml", cfg)
	if err != nil {
		t.Fatalf("Can't load config file. Error %s", err.Error())
	}

	for i, proxy := range cfg.Proxies {
		if proxy.Url != fmt.Sprintf("test_%d_url", (i+1)) {
			t.Fatalf("Wrong Url from config %s", proxy.Url)
		}

		if proxy.Endpoint != fmt.Sprintf("test_%d_endpoint", (i+1)) {
			t.Fatalf("Wrong Endpoint from config %s", proxy.Endpoint)
		}
	}
}

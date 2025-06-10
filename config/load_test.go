package config_test

import (
	"Book-API-Server/config"
	"testing"
)

func TestLoadConfigFromYaml(t *testing.T) {
	if err := config.LoadConfigFromYaml("application.yaml"); err != nil {
		t.Fatal(err)
	}
	t.Log(config.Get())
}

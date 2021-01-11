package Config

import (
	"os"
	"reflect"
	"testing"
)

var testUser = os.Getenv("LOGIN")
var testPass = os.Getenv("PASSWORD")
var yamlConfig = YamlConfig{testUser, "", configPath}

func TestCreateConfig(t *testing.T) {
	if err := yamlConfig.CreateConfig(); err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Error(err)
	}
}

func TestUpdateConfig(t *testing.T) {
	yamlConfig.Token = "TestToken"
	if err := yamlConfig.UpdateConfig(); err != nil {
		t.Error(err)
	}
}

func TestLoadToken(t *testing.T) {
	token, err := LoadToken()
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Errorf("Token in config is empty")
	}
}

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig()
	if err != nil {
		t.Error(err)
	}
	v := reflect.ValueOf(config)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			t.Errorf("Empty Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		}
	}
}

func deleteConfig(t *testing.T) {
	if err := os.Remove(configPath); err != nil {
		t.Error(err)
	}
}

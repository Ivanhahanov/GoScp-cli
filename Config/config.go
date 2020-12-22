package Config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)
var configPath = os.Getenv("HOME") + "/.scp.yaml"

type YamlConfig struct {
	Username string `yaml:"username"`
	Token string `yaml:"token"`
	RootDir string `yaml:"rootDir"`
}

func (c *YamlConfig)UpdateConfig() error{
	filename, _ := filepath.Abs(configPath)
	data, err := yaml.Marshal(&c)
	if err != nil{
		return err
	}
	if err = ioutil.WriteFile(filename, data, 0644); err != nil{
		return err
	}
	return nil
}

func (c *YamlConfig)CreateConfig() error{
	filename, _ := filepath.Abs(configPath)
	f, err := os.Create(filename)
	if err != nil{
		return err
	}
	defer f.Close()
	return nil
}

func UpdateToken(token string) error {
	var c YamlConfig

	filename, _ := filepath.Abs(configPath)
	file, err := ioutil.ReadFile(filename)
	if err != nil{
		return err
	}
	if err = yaml.Unmarshal(file, &c); err != nil{
		return err
	}
	c.Token = token
	if err = c.UpdateConfig(); err != nil{
		return err
	}
	return nil
}

func LoadToken() (token string, err error){
	var c YamlConfig
	filename, _ := filepath.Abs(configPath)
	file, err := ioutil.ReadFile(filename)
	if err != nil{
		return "", err
	}
	if err = yaml.Unmarshal(file, &c); err != nil{
		return "", err
	}
	return c.Token, err
}

func LoadConfig() (c YamlConfig, err error){
	filename, _ := filepath.Abs(configPath)
	file, err := ioutil.ReadFile(filename)
	if err != nil{
		return c, err
	}
	if err = yaml.Unmarshal(file, &c); err != nil{
		return c, err
	}
	return c, err
}
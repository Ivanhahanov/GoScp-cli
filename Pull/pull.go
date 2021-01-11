package Pull

import (
	"SecureCodePlatform-cli/Config"
	"SecureCodePlatform-cli/Login"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ChallengeConfig struct {
	ChallengeId string `yaml:"challenge_id"`
}

type ChallengeName struct {
	Name   string `json:"title"`
	Detail string `json:"detail"`
}

func (c *ChallengeConfig) CreateChallengeConfig(path string) error {
	filename := filepath.Join(path, "config.yml")
	data, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, data, 0644); err != nil {
		return err
	}
	return nil
}

func GetChallenge(id string) (err error) {
	path, err := CreateChallengeDir(id)
	if err != nil {
		return err
	}
	challengeConfig := &ChallengeConfig{}
	challengeConfig.ChallengeId = id
	if err = challengeConfig.CreateChallengeConfig(path); err != nil {
		return err
	}
	return nil
}

func CreateChallengeDir(id string) (path string, err error) {
	config, err := Config.LoadConfig()
	if err != nil {
		return "", err
	}
	dirName, err := GetChallengeName(id)
	fmt.Println(dirName)
	path = filepath.Join(config.RootDir, dirName)
	fmt.Println(path)
	if err = os.Mkdir(path, os.ModePerm); err != nil {
		return "", err
	}
	return path, nil
}

func GetChallengeName(id string) (name string, err error) {
	token, _ := Config.LoadToken()
	client := http.Client{Timeout: time.Second * 2}
	var bearer = "Bearer " + token
	req, _ := http.NewRequest("GET", Login.DefaultUrl+"/challenges/get_challenge", nil)
	req.Header.Add("Authorization", bearer)
	q := req.URL.Query()
	q.Add("challenge_id", id)
	req.URL.RawQuery = q.Encode()

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", readErr
	}
	fmt.Println(string(body))
	challengeName := ChallengeName{}
	jsonErr := json.Unmarshal(body, &challengeName)
	if jsonErr != nil {
		return challengeName.Name, jsonErr
	}

	return strings.ReplaceAll(challengeName.Name, " ", ""), nil
}

func GetFiles(id string) error {
	return nil
}

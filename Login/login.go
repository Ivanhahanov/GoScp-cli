package Login

import (
	"SecureCodePlatform-cli/Config"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var DefaultUrl = "http://localhost"

type Credentials struct {
	Username string
	Password string
}

type Answer struct {
	Token  string `json:"access_token"`
	User   string `json:"username"`
	Detail string `json:"detail"`
}

func (c *Credentials) CheckCredentials() (token string, err error) {
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	data := url.Values{
		"password": {c.Password},
		"username": {c.Username},
	}
	req, _ := http.NewRequest("POST", DefaultUrl+"/users/token", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", readErr
	}
	answer, err := GetData(body)
	if err != nil {
		return "", err
	}

	if answer.Detail == "" {
		return answer.Token, nil
	} else {
		return "", fmt.Errorf(answer.Detail)
	}
}

func EnterLoginAndPass() (username string, err error) {
	fmt.Print("Enter Username: ")
	if _, err := fmt.Scanf("%s", &username); err != nil {
		return "", err
	}

	if _, err := EnterPass(username); err != nil {
		return "", err
	}
	return username, nil
}
func EnterPass(username string) (token string, err error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}
	fmt.Println()
	password := string(bytePassword)
	if token, err = Auth(username, password); err != nil {
		return "", err
	}
	return token, nil
}

func Auth(username string, password string) (token string, err error) {
	cred := Credentials{}
	cred.Username = username
	cred.Password = password
	if token, err = cred.CheckCredentials(); err != nil {
		return "", err
	}
	config := Config.YamlConfig{
		Username: username,
		Token:    token,
		RootDir:  os.Getenv("HOME"),
	}
	if err := config.UpdateConfig(); err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken() (token string, err error) {
	c, err := Config.LoadConfig()
	if err != nil {
		return "", err
	}
	token = c.Token
	client := http.Client{Timeout: time.Second * 2}
	var bearer = "Bearer " + token
	req, _ := http.NewRequest("GET", DefaultUrl+"/users/me", nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", readErr
	}
	if _, err = GetData(body); err != nil {
		if token, err := EnterPass(c.Username); err == nil {
			return token, err
		}
		return "", err
	}

	return token, nil
}

func GetData(data []byte) (answer Answer, err error) {
	answer = Answer{}
	jsonErr := json.Unmarshal(data, &answer)
	if jsonErr != nil {
		return answer, jsonErr
	}
	if answer.Detail != "" {
		return answer, fmt.Errorf(answer.Detail)
	}
	return answer, nil
}

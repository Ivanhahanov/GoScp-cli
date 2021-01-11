package Login

import (
	"SecureCodePlatform-cli/Config"
	"fmt"
	"os"
	"testing"
)

var testUser = os.Getenv("LOGIN")
var testPass = os.Getenv("PASSWORD")

func TestLoginAuth(t *testing.T) {
	token, err := Auth(testUser, testPass)
	if err != nil {
		t.Error(err)
	}
	if err := checkTokenConfig(token); err != nil {
		t.Error(err)
	}
}

type LoginPass struct {
	Login string
	Pass  string
}

var badCredential = []LoginPass{
	{"admin", "admin"},
	{"user", " "},
	{"user", "password"},
}

func TestBadCredentials(t *testing.T) {
	for _, cred := range badCredential {
		_, err := Auth(cred.Login, cred.Pass)
		if fmt.Sprint(err) != "Incorrect username or password" {
			t.Error(err)
		}
	}

}

func checkTokenConfig(token string) error {
	c, err := Config.LoadConfig()
	if err != nil {
		return err
	}
	if c.Token != token {
		return fmt.Errorf("Config Token != Auth token")
	}
	return nil
}

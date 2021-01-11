package Score

import (
	"SecureCodePlatform-cli/Login"
	"os"
	"reflect"
	"testing"
)

var testUser = os.Getenv("LOGIN")
var testPass = os.Getenv("PASSWORD")

func TestGetScoreboard(t *testing.T) {
	token, err := Login.Auth(testUser, testPass)
	if err != nil {
		t.Error(err)
	}
	score, err := GetScoreboard(token)
	if err != nil {
		t.Error(err)
	}
	v := reflect.ValueOf(*score)
	typeOfS := v.Type()
	if score.Detail != "" {
		t.Errorf(score.Detail)
	}
	for i := 0; i < v.NumField()-1; i++ {
		if v.Field(i).Interface() == "" {
			t.Errorf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		}
	}
}

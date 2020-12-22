package Score

import (
	"SecureCodePlatform-cli/Login"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Scoreboard struct {
	Username string `json:"username"`
	UsersScore int `json:"users_score"`
	UsersPlace int `json:"users_place"`
	NumOfUsers int `json:"num_of_users"`
	NumOfChallenges int `json:"num_of_challenges"`
	NumOfSolvedChallenges int `json:"num_of_solved_challenges"`
	Detail string `json:"detail"`
}

func GetScoreboard(token string) (scoreboard *Scoreboard, err error){
	client := http.Client{Timeout: time.Second * 2}
	var bearer = "Bearer " + token
	req, _ := http.NewRequest("GET", Login.DefaultUrl+"/scoreboard/info", nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return scoreboard, readErr
	}
	jsonErr := json.Unmarshal(body, &scoreboard)
	if jsonErr != nil {
		return scoreboard, jsonErr
	}
	if scoreboard.Detail != "" {
		return scoreboard, fmt.Errorf(scoreboard.Detail)
	} else {
		return scoreboard, nil
	}

}

func PrintScoreboard(token string) error {
	score, err := GetScoreboard(token)
	if err != nil{
		return err
	}
	fmt.Printf("Username:   %s\n", score.Username)
	fmt.Printf("Score:      %d\n", score.UsersScore)
	fmt.Printf("Challenges: %d/%d\n", score.NumOfSolvedChallenges, score.NumOfChallenges)
	fmt.Printf("Place:      %d/%d\n", score.UsersPlace, score.NumOfUsers)
	return nil
}
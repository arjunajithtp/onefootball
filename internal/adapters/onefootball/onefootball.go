package ofb

import (
	"github.com/arjunajithtp/onefootball/public/dtos/api/ofbapi"
	"net/http"
	"fmt"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"log"
)

type OneFootBall interface {
	GetTeamDetails(teamID int) (*ofbapi.OneFootBallTeamDetails, error)
}

type Downloader struct{}

func (d *Downloader) GetTeamDetails(teamID int) (*ofbapi.OneFootBallTeamDetails, error) {
	var response ofbapi.OneFootBallTeamDetails
	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodGet,
		"https://vintagemonster.onefootball.com/api/teams/en/" + strconv.Itoa(teamID) + ".json",
		nil,
		)
	if err != nil {
		return nil, fmt.Errorf("error while creating GET Method request for ofb: %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while processing POST Method request for ofb: %v", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("info: recieved invalid response from ofb with team id: ", teamID)
		return nil, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading GET response body from ofb: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error while trying to unmarshal ofb response: %v, %v", err, teamID)
	}
	return &response, nil
}
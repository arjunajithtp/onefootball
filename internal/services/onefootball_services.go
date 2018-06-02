package services

import (
	"fmt"
	"github.com/arjunajithtp/onefootball/internal/adapters/onefootball"
	"log"
)

type Downloader struct {
	ofbDownloader ofb.OneFootBall
}

type TeamDetails struct {
	Name             string
	AllPlayerDetails []Player
}

type Player struct {
	Age      string
	TeamList string
}

func GetPlayerDetails(in ofb.OneFootBall, requiredTeams map[string]bool) (map[string]*Player, []string, error) {

	var d = Downloader{ofbDownloader: in}
	allPlayers, playerNames, err := d.GetAllPlayers(requiredTeams)
	if err != nil {
		return nil, nil, fmt.Errorf("error while trying to get team details: %v", err)
	}

	return allPlayers, playerNames, nil

}

func (d *Downloader) GetAllPlayers(requiredTeams map[string]bool) (map[string]*Player, []string, error) {

	playersMap := make(map[string]*Player)
	var playerNames []string
	flag := 1
	for i := 0; ; i++ {
		teamData, err := d.ofbDownloader.GetTeamDetails(i)
		if err != nil {
			return nil, nil, fmt.Errorf("error while trying to get team details from ofb: %v", err)
		}

		if teamData == nil || !requiredTeams[teamData.Data.Team.Name] {
			continue
		}

		log.Printf("Found details of team %v with team_id %v...", teamData.Data.Team.Name, i)
		for _, teamPlayer := range teamData.Data.Team.Players {
			fullName := teamPlayer.FirstName + " " + teamPlayer.LastName
			if playersMap[fullName] != nil {
				playersMap[fullName].TeamList += ", " + teamData.Data.Team.Name
				continue
			}
			playersMap[fullName] = &Player{
				Age:      teamPlayer.Age,
				TeamList: teamData.Data.Team.Name,
			}
			playerNames = append(playerNames, fullName)
		}

		if flag == len(requiredTeams) {
			return playersMap, playerNames, nil
		}
		flag++

	}
	return playersMap, playerNames, nil
}

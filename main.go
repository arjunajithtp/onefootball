package main

import (
	"github.com/arjunajithtp/onefootball/internal/adapters/onefootball"
	"github.com/arjunajithtp/onefootball/internal/services"
	"log"
	"sort"
	"fmt"
)

func main() {
	ofbDownloader := ofb.Downloader{}

	playerDertails, playerNames, err := services.GetPlayerDetails(&ofbDownloader, map[string]bool{
		"Germany": true,
		"England": true,
		"France": true,
		"Spain": true,
		"Manchester Utd": true,
		"Arsenal": true,
		"Chelsea": true,
		"Barcelona": true,
		"Real Madrid": true,
		"FC Bayern Munich": true,
	})
	if err != nil {
		log.Println("error while trying to get player details for output: ", err)
	}
	log.Println("Player details are as follows: ")
	sort.Strings(playerNames)

	log.Println(playerNames)
	for i := 0; i < len(playerNames); i++ {
		fmt.Printf("%v; %v; %v \n", playerNames[i], playerDertails[playerNames[i]].Age, playerDertails[playerNames[i]].TeamList)
	}


}
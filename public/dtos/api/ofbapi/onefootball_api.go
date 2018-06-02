package ofbapi

type OneFootBallTeamDetails struct {
	Data	Data	`json:"data"`
}

type Data struct {
	Team	Team	`json:"team"`
}

type Team struct {
	Name	string	`json:"name"`
	Players	[]Player	`json:"players"`
}

type Player struct {
	FirstName	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
	Age	string	`json:"age"`
}
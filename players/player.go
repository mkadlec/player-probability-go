package players

type Player struct {
	Name        string
	Probability int
}

func All() []Player {
	var playerProbabilities []Player
	playerProbabilities = append(playerProbabilities, Player{"Dave", 90})
	playerProbabilities = append(playerProbabilities, Player{"Kevin", 90})
	playerProbabilities = append(playerProbabilities, Player{"Simon", 90})
	playerProbabilities = append(playerProbabilities, Player{"Kenny", 90})
	playerProbabilities = append(playerProbabilities, Player{"Gary", 90})
	playerProbabilities = append(playerProbabilities, Player{"Cody", 90})
	playerProbabilities = append(playerProbabilities, Player{"Danny", 90})
	playerProbabilities = append(playerProbabilities, Player{"Mike", 90})
	playerProbabilities = append(playerProbabilities, Player{"Pete", 90})
	playerProbabilities = append(playerProbabilities, Player{"Shaun", 90})
	playerProbabilities = append(playerProbabilities, Player{"Sato", 60})
	return playerProbabilities
}

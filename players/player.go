package players

type Player struct {
	Name        string
	Probability int
}

func All() []Player {
	var playerProbabilities []Player
	playerProbabilities = append(playerProbabilities, Player{"Jeff", 90})
	playerProbabilities = append(playerProbabilities, Player{"Shaun", 90})
	return playerProbabilities
}

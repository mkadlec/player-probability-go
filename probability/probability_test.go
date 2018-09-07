package probability

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/mkadlec/player-probability-go/players"

func TestProbability(t *testing.T) {
	var playerProbabilities []players.Player
	playerProbabilities = append(playerProbabilities, players.Player{"Test", 100})

	p := Calculator{PlayerProbabilities: playerProbabilities}
	var player_count = p.HowManyPlayers()

	assert.Equal(t, player_count, 2, "they should be equal")

	var playerProbabilitiesNone []players.Player
	playerProbabilitiesNone = append(playerProbabilitiesNone, players.Player{"Test", 0})

	p2 := Calculator{PlayerProbabilities: playerProbabilitiesNone}
	var players2 = p2.HowManyPlayers()

	assert.Equal(t, players2, 1, "they should be equal")
}

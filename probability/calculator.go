package probability

import (
	"math"
	"math/rand"
	"time"

	"github.com/mkadlec/player-probability-go/players"
)

var MYSELF = 1

type Calculator struct {
	PlayerProbabilities []players.Player
}

func (Calculator *Calculator) HowManyPlayers() int {
	rand.Seed(time.Now().UnixNano())
	iterations := 10
	totalGuests := 0

	var probabilities = Calculator.PlayerProbabilities
	var length = len(probabilities)

	for i := 1; i <= iterations; i++ {
		var players = 0
		for playerIndex := 0; playerIndex < length; playerIndex++ {
			var randomDiceThrow = rand.Intn(100)
			var playerProbability = probabilities[playerIndex].Probability
			if randomDiceThrow < playerProbability {
				players += 1
			}
		}
		totalGuests += players
	}

	var guests = math.Round((float64(totalGuests) / float64(iterations)))

	return MYSELF + int(guests)
}

package probability

import (
	"fmt"
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
			fmt.Printf("dice says %v, has to beat %v \n", randomDiceThrow, playerProbability)
			if randomDiceThrow < playerProbability {
				players += 1
				fmt.Printf("Adding, players now %v\n", players)
			}
		}
		totalGuests += players
		fmt.Printf("Total Guests %v\n", totalGuests)
	}

	var guests = math.Round((float64(totalGuests) / float64(iterations)))
	fmt.Printf("Final calc %v / %v = %v\n", totalGuests, iterations, guests)

	return MYSELF + int(guests)
}

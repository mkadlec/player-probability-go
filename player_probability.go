package main

import "fmt"
import "github.com/mkadlec/player-probability-go/probability"
import "github.com/mkadlec/player-probability-go/players"

func main() {
	var playerProbabilities = players.All()
	p := probability.Calculator{PlayerProbabilities: playerProbabilities}
	fmt.Printf("%v players will show\n", p.HowManyPlayers())
}

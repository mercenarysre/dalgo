package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type player struct {
	Name       string `json:"name"`
	High_Score int    `json:"high_score"`
}

var players []player

func jsonParser() {
	playerData, err := os.ReadFile("json.txt")
	if err != nil {
		fmt.Println("failed to open file", err)
	}

	err = json.Unmarshal(playerData, &players)
	if err != nil {
		fmt.Println("failed to unmarshal JSON data", err)
	}
	fmt.Println(players)

	l := len(players)
	highestScorer := players[0].High_Score
	lowestScorer := players[0].High_Score

	for i := 1; i < l; i++ {
		if players[i].High_Score >= highestScorer {
			highestScorer = players[i].High_Score
			fmt.Println("Player with highest score: ", players[i].Name)
		}

		if players[i].High_Score <= lowestScorer {
			lowestScorer = players[i].High_Score
			fmt.Println("Player with lowest score: ", players[i].Name)
		}

	}

}

func main() {
	jsonParser()
}

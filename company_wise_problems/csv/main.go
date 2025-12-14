package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	name       string
	high_score int
}

func csvParser() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("failed to open file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("failed to read records from the file:", err)
	}

	var players []Player
	for _, record := range records[1:] {
		high_score, _ := strconv.Atoi(record[1])
		player := Player{
			name:       record[0],
			high_score: high_score,
		}
		players = append(players, player)
	}

	fmt.Println(players)

	l := len(players)
	highestScorer := players[0].high_score
	lowestScorer := players[0].high_score

	for i := 1; i < l; i++ {
		if players[i].high_score >= highestScorer {
			highestScorer = players[i].high_score
			fmt.Println("Player with highest score: ", players[i].name)
		}

		if players[i].high_score <= lowestScorer {
			lowestScorer = players[i].high_score
			fmt.Println("Player with lowest score: ", players[i].name)
		}

	}
}

func main() {
	csvParser()
}

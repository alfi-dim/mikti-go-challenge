package main

import (
	"fmt"
	"math"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func PickAWinner(team1 Team, team2 Team) {

	if team1.Avg > team2.Avg {
		fmt.Printf("Tim %s menjadi pemenang karena memiliki skor lebih tinggi dari tim lawan\n", team1.Name)
		var isScoreMoreThan100 bool
		for _, score := range team1.Scores {
			if score > 100 {
				isScoreMoreThan100 = true
				break
			}
		}
		if isScoreMoreThan100 {
			fmt.Println("dan memiliki skore lebih dari 100")
		}
	} else if team1.Avg < team2.Avg {
		fmt.Printf("Tim %s menjadi pemenang karena memiliki skor lebih tinggi dari tim lawan\n", team2.Name)
		var isScoreMoreThan100 bool
		for _, score := range team2.Scores {
			if score > 100 {
				isScoreMoreThan100 = true
				break
			}
		}
		if isScoreMoreThan100 {
			fmt.Println("dan memiliki skore lebih dari 100")
		}
	} else {
		var isTeam1ScoreMoreThan100 bool
		var isTeam2ScoreMoreThan100 bool

		for _, score := range team1.Scores {
			if score > 100 {
				isTeam1ScoreMoreThan100 = true
				break
			}
		}
		for _, score := range team2.Scores {
			if score > 100 {
				isTeam2ScoreMoreThan100 = true
				break
			}
		}

		if isTeam1ScoreMoreThan100 && isTeam2ScoreMoreThan100 {
			fmt.Println("Hasil seri, kedua tim memiliki skor lebih dari 100")
			fmt.Println("Kedua tim mendapat trofi")
		} else {
			fmt.Println("Hasil seri, kedua tim memiliki skor kurang dari 100")
			fmt.Println("Jadi tidak ada yang mendapat trofi")
		}
	}
}

func (team Team) CalculateScoreAverage() float32 {
	var averageScore float32
	var totalScores float32
	for _, score := range team.Scores {
		totalScores += float32(score)
	}
	averageScore = totalScores / float32(len(team.Scores))
	parsedFloat := roundFloat(float64(averageScore), 2)
	return float32(parsedFloat)
}

package main

import (
	"fmt"
	"math"
)

// fungsi helper untuk membulatkan dua angka di belakang koma
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// fungsi untuk memilih pemenang
func PickAWinner(team1 Team, team2 Team) {
	// membuat variabel untuk menampung data yang diperlukan
	var winner Team                   // variabel untuk menampung tim pemenang
	var isWinnerScoreMoreThan100 bool // variabel untuk menampung informasi apakah tim pemenang memiliki skor lebih dari 100
	var isDraw bool                   // variabel untuk menampung informasi apakah hasilnya seri

	// membandingkan rata-rata skor dari masing-masing tim
	if team1.Avg > team2.Avg {
		winner = team1
	} else if team1.Avg < team2.Avg {
		winner = team2
	} else {
		isDraw = true
	}

	// jika hasilnya seri
	if isDraw {
		// membuat variabel untuk menampung informasi apakah kedua tim memiliki skor lebih dari 100
		var isTeam1ScoreMoreThan100 bool
		var isTeam2ScoreMoreThan100 bool

		// memeriksa apakah kedua tim memiliki skor lebih dari 100
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

		// jika kedua tim memiliki skor lebih dari 100
		if isTeam1ScoreMoreThan100 && isTeam2ScoreMoreThan100 {
			fmt.Println("Hasil seri, kedua tim memiliki skor lebih dari 100")
			fmt.Println("Kedua tim mendapat trofi")
		} else { // jika kedua tim memiliki skor kurang dari 100
			fmt.Println("Hasil seri, kedua tim memiliki skor kurang dari 100")
			fmt.Println("Jadi tidak ada yang mendapat trofi")
		}
	} else { // jika hasilnya tidak seri, atau ada pemenang di antara kedua tim
		fmt.Printf("Tim %s menjadi pemenang karena memiliki skor lebih tinggi dari tim lawan\n", winner.Name)

		// memeriksa apakah pemenang memiliki skor lebih dari 100
		for _, score := range winner.Scores {
			if score > 100 {
				isWinnerScoreMoreThan100 = true
				break
			}
		}

		// jika pemenang memiliki skor lebih dari 100
		if isWinnerScoreMoreThan100 {
			fmt.Println("dan memiliki skore lebih dari 100")
		}
	}
}

// fungsi untuk menghitung rata-rata skor
func (team Team) CalculateScoreAverage() float32 {
	// membuat variabel untuk menampung total skor dan rata-rata skor
	var averageScore float32
	var totalScores float32

	// menghitung total skor
	for _, score := range team.Scores {
		totalScores += float32(score)
	}

	// menghitung rata-rata skor
	averageScore = totalScores / float32(len(team.Scores))

	// membulatkan rata-rata skor menjadi dua angka di belakang koma
	parsedFloat := roundFloat(float64(averageScore), 2)
	return float32(parsedFloat)
}

// fungsi untuk menghitung BMI
func (person Person) CalculateBMIScore() float32 {
	bmi := person.Mass / (person.Height * person.Height)

	parsedFloat := roundFloat(float64(bmi), 2)
	return float32(parsedFloat)
}

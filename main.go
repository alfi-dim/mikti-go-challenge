package main

import "fmt"

func main() {
	Soal1()
}

func Soal1() {
	LumbaLumba := Team{
		Name: "Lumba-Lumba",
	}
	Koala := Team{
		Name: "Koala",
	}

	fmt.Println("Masukkan Skor untuk Tim Lumba-Lumba")
	for i := 1; i <= 3; i++ {
		var score int
		fmt.Printf("Skor ke: %d untuk Tim Lumba-Lumba: ", i)
		_, err := fmt.Scan(&score)
		if err != nil {
			return
		}
		LumbaLumba.Scores = append(LumbaLumba.Scores, Score(score))
	}
	fmt.Println()
	fmt.Println("Masukkan Skor untuk Tim Koala")
	for i := 1; i <= 3; i++ {
		var score int
		fmt.Printf("Masukkan Skor ke: %d untuk Tim Koala: ", i)
		_, err := fmt.Scan(&score)
		if err != nil {
			return
		}
		Koala.Scores = append(Koala.Scores, Score(score))
	}
	LumbaLumba.Avg = LumbaLumba.CalculateScoreAverage()
	Koala.Avg = Koala.CalculateScoreAverage()
	PickAWinner(LumbaLumba, Koala)
}

package challenge1

import "fmt"

func main() {
	LumbaLumba := Team{
		Name: "Lumba-Lumba",
	}
	Koala := Team{
		Name: "Koala",
	}

	for i := 1; i <= 3; i++ {
		var score int
		fmt.Printf("Masukkan Skor ke: %d untuk Tim Lumba-Lumba: \n", i)
		_, err := fmt.Scan(&score)
		if err != nil {
			return
		}
		LumbaLumba.Scores = append(LumbaLumba.Scores, Score(score))
	}

	for i := 1; i <= 3; i++ {
		var score int
		fmt.Printf("Masukkan Skor ke: %d untuk Tim Koala: \n", i)
		_, err := fmt.Scan(&score)
		if err != nil {
			return
		}
		Koala.Scores = append(Koala.Scores, Score(score))
	}

	fmt.Println(LumbaLumba)
	fmt.Println(Koala)
}

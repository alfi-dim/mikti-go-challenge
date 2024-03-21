package main

import "fmt"

func main() {
	Soal1()
	Soal2()
}

// fungsi untuk soal 1
func Soal1() {
	// menambahkan tim
	LumbaLumba := Team{
		Name: "Lumba-Lumba",
	}
	Koala := Team{
		Name: "Koala",
	}

	// memasukkan skor untuk masing-masing tim
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

	// menghitung rata-rata skor untuk masing-masing tim
	LumbaLumba.Avg = LumbaLumba.CalculateScoreAverage()
	Koala.Avg = Koala.CalculateScoreAverage()

	// memilih pemenang
	PickAWinner(LumbaLumba, Koala)
}

func Soal2() {
	Mark := Person{
		Name: "Mark",
	}
	John := Person{
		Name: "John",
	}
	var markHigherBMI bool

	fmt.Print("Masukkan Berat Badan Mark: ")
	_, err := fmt.Scan(&Mark.Mass)
	if err != nil {
		return
	}
	fmt.Print("Masukkan Tinggi Badan Mark: ")
	_, err = fmt.Scan(&Mark.Height)
	if err != nil {
		return
	}

	Mark.BMI = Mark.CalculateBMIScore()

	fmt.Print("Masukkan Berat Badan John: ")
	_, err = fmt.Scan(&John.Mass)
	if err != nil {
		return
	}
	fmt.Print("Masukkan Tinggi Badan John: ")
	_, err = fmt.Scan(&John.Height)
	if err != nil {
		return
	}

	John.BMI = John.CalculateBMIScore()

	if Mark.BMI > John.BMI {
		markHigherBMI = true
	} else {
		markHigherBMI = false
	}

	if markHigherBMI {
		fmt.Println("BMI Mark lebih tinggi dari BMI John")
	} else {
		fmt.Println("BMI John lebih tinggi dari BMI Mark")
	}
}

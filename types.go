package main

// mendeklarasikan type Score dengan tipe data float32
type Score float32

// mendeklarasikan type Team dengan tipe data struct
type Team struct {
	Name   string
	Scores []Score
	Avg    float32
}

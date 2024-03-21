package main

type Score float32
type Team struct {
	Name   string
	Scores []Score
	Avg    float32
}

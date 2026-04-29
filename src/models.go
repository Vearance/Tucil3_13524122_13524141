package main

type Point struct {
	X, Y int
}

type Board struct {
	N, M   int
	Grid   [][]rune
	Costs  [][]int
}

type State struct {
	Pos           Point
	Path          string
	TotalCost     int
	Steps         int
	CurrentTarget int  // untuk urutan angka 0-9
}
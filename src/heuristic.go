package main

import "math"

// Heuristic 1: Manhattan Distance
// h(n) = |xi - xp| + |yi - yp|
func ManhattanDistance(pos Point, board *Board, target int) int {
	var targetPoint Point
	found := false

	// jika target masih ada di papan (kurang dari atau sama dengan angka tertinggi)
	if target <= board.MaxTarget && board.MaxTarget != -1 {
		targetChar := rune(target + '0')
		for i := 0; i < board.N; i++ {
			for j := 0; j < board.M; j++ {
				if board.Grid[i][j] == targetChar {
					targetPoint = Point{i, j}
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	} 
	
	// jika semua angka sudah diambil, atau memang tidak ada angka di papan
	if !found {
		for i := 0; i < board.N; i++ {
			for j := 0; j < board.M; j++ {
				if board.Grid[i][j] == 'O' {
					targetPoint = Point{i, j}
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	return int(math.Abs(float64(pos.X-targetPoint.X)) + math.Abs(float64(pos.Y-targetPoint.Y)))
}
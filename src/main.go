package main

import (
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	var filename string
	fmt.Print("Masukan file input : ")
	fmt.Scanln(&filename)
	filename = filepath.Join("test", filepath.Base(filename))

	var algo string
	fmt.Print("Algoritma apa yang anda pilih? (UCS/GBFS/A*) ")
	fmt.Scanln(&algo)

	var heuristic string
	if algo != "UCS" {
		fmt.Print("Heuristic apa yang anda pilih? (H1/H2) ")
		fmt.Scanln(&heuristic)
	}

	board, startPos, err := LoadMap(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var hFunc func(Point, *Board, int) int
	switch heuristic {
		case "H1":
			hFunc = ManhattanDistance
		case "H2":
			hFunc = EuclideanDistance
	}

	start := time.Now()
	var r res
	switch algo {
	case "UCS":
		r = UCS(board, startPos)
	case "GBFS":
		if hFunc == nil {
			fmt.Println("Heuristik tidak valid")
			return
		}
		r = GBFS(board, startPos, hFunc)
	case "A*":
		if hFunc == nil {
			fmt.Println("Heuristik tidak valid")
			return
		}
		r = AStar(board, startPos, hFunc)
	default:
		fmt.Println("Algoritma tidak valid")
		return
	}
	elapsed := time.Since(start)

	if !r.Found {
		fmt.Println("Solusi tidak ditemukan")
	} else {
		fmt.Printf("Solusi Yang Ditemukan : %s\n", r.Path)
		fmt.Printf("Cost dari Solusi : %d\n", r.Cost)
	}
	fmt.Printf("Waktu eksekusi: %d ms\n", elapsed.Milliseconds())
	fmt.Printf("Banyak iterasi yang dilakukan: %d iterasi\n", r.Iter)
}
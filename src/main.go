package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var filename string
	fmt.Print("Masukan file input : ")
	fmt.Scanln(&filename)
	filename = filepath.Join("test", filepath.Base(filename))

	board, startPos, err := LoadMap(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var algo string
	fmt.Print("Algoritma apa yang anda pilih? (UCS/GBFS/A*) ")
	fmt.Scanln(&algo)
	algo = strings.ToUpper(algo)

	var heuristic string
	if algo != "UCS" {
		fmt.Print("Heuristic apa yang anda pilih? (H1/H2) ")
		fmt.Scanln(&heuristic)
		heuristic = strings.ToUpper(heuristic)
	}


	var hFunc func(Point, *Board, int) int
	switch heuristic {
		case "H1":
			hFunc = ManhattanDistance
		case "H2":
			hFunc = EuclideanDistance
	}

	start := time.Now()
	var goal *State
	var iter int
	switch algo {
	case "UCS":
		goal, iter = UCS(board, startPos)
	case "GBFS":
		if hFunc == nil {
			fmt.Println("Heuristik tidak valid")
			return
		}
		goal, iter = GBFS(board, startPos, hFunc)
	case "A*":
		if hFunc == nil {
			fmt.Println("Heuristik tidak valid")
			return
		}
		goal, iter = AStar(board, startPos, hFunc)
	default:
		fmt.Println("Algoritma tidak valid")
		return
	}
	dT := time.Since(start)

	if goal == nil {
		fmt.Println("Solusi tidak ditemukan")
		fmt.Printf(">> Waktu eksekusi: %d ms\n", dT.Milliseconds())
		fmt.Printf(">> Banyak iterasi yang dilakukan: %d iterasi\n", iter)
		return
	}

	path, pos := reverse(goal)
	cost := goal.TotalCost

	printSolution(board, startPos, path, pos, cost)

	fmt.Printf("\n>> Waktu eksekusi: %d ms\n", dT.Milliseconds())
	fmt.Printf(">> Banyak iterasi yang dilakukan: %d iterasi\n", iter)

	var pb string
	fmt.Print(">> Apakah Anda ingin melakukan playback? (Ya/Tidak) : ")
	fmt.Scanln(&pb)
	if strings.ToUpper(pb) == "YA" {
		step := 0
		quit := false
		for !quit {
			fmt.Println()
			if step == 0 {
				fmt.Println("Initial")
			} else {
				fmt.Printf("Step %d : %c\n", step, path[step-1])
			}
			fmt.Print(drawBoard(board, startPos, pos[step]))
			fmt.Print("[A=prev D=next J=jump Q=quit]: ")
			var cmd string
			fmt.Scanln(&cmd)
			switch cmd {
			case "A", "a":
				if step > 0 {
					step--
				}
			case "D", "d":
				if step < len(pos)-1 {
					step++
				}
			case "J", "j":
				fmt.Print("Step: ")
				var n int
				fmt.Scanln(&n)
				if n >= 0 && n < len(pos) {
					step = n
				}
			case "Q", "q":
				quit = true
			}
		}
	}

	var sv string
	fmt.Print(">> Apakah Anda ingin menyimpan solusi? (Ya/Tidak) : ")
	fmt.Scanln(&sv)
	if strings.ToUpper(sv) == "YA" {
		var fn string
		fmt.Print(">> Masukan nama file untuk menyimpan solusi : ")
		fmt.Scanln(&fn)
		if fn == "" {
			fn = "solusi.txt"
		}
		if errSv := saveSolution(fn, board, startPos, path, pos, cost, iter, dT.Milliseconds()); errSv != nil {
			fmt.Println("Gagal menyimpan:", errSv)
		} else {
			abs, _ := filepath.Abs(fn)
			fmt.Printf(">> Solusi disimpan pada %s\n", abs)
		}
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func scanToken(reader *bufio.Reader) string {
	var token string
	fmt.Fscan(reader, &token)
	return token
}

func choice(reader *bufio.Reader, prompt string, valid map[string]bool) string {
	for {
		fmt.Print(prompt)
		ch := strings.ToUpper(scanToken(reader))
		if valid[ch] {
			return ch
		}
		fmt.Println("Input tidak valid")
	}
}

func resolveMapPath(input string) string {
	if input == "" {
		return ""
	}
	if _, err := os.Stat(input); err == nil {
		return input
	}
	fallback := filepath.Join("test", filepath.Base(input))
	if _, err := os.Stat(fallback); err == nil {
		return fallback
	}
	return fallback
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var (
		board    *Board
		startPos Point
		err      error
	)

	for {
		fmt.Print("Masukan file input : ")
		inputFile := scanToken(reader)
		filename := resolveMapPath(inputFile)
		board, startPos, err = LoadMap(filename)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	algo := choice(reader, "Algoritma apa yang anda pilih? (UCS/GBFS/A*) ", map[string]bool{
		"UCS":  true,
		"GBFS": true,
		"A*":   true,
	})

	heuristic := ""
	if algo != "UCS" {
		heuristic = choice(reader, "Heuristic apa yang anda pilih? (H1/H2/H3) ", map[string]bool{
			"H1": true,
			"H2": true,
			"H3": true,
		})
	}

	var hFunc func(Point, *Board, int) int
	switch heuristic {
		case "H1":
			hFunc = ManhattanDistance
		case "H2":
			hFunc = EuclideanDistance
		case "H3":
			hFunc = ChebyshevDistance
	}

	start := time.Now()
	var goal *State
	var iter int
	switch algo {
	case "UCS":
		goal, iter = UCS(board, startPos)
	case "GBFS":
		goal, iter = GBFS(board, startPos, hFunc)
	case "A*":
		goal, iter = AStar(board, startPos, hFunc)
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

	for {
		fmt.Print(">> Apakah Anda ingin melakukan playback? (Y/N) : ")
		pb := strings.ToUpper(scanToken(reader))
		if pb == "N" {
			break
		}
		if pb != "Y" {
			fmt.Println("Input tidak valid")
			continue
		}

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
			cmd := strings.ToUpper(scanToken(reader))
			switch cmd {
			case "A":
				if step > 0 {
					step--
				} else {
					fmt.Println(">> Posisi utama")
				}
			case "D":
				if step < len(pos)-1 {
					step++
				} else {
					fmt.Println(">> Posisi akhir")
				}
			case "J":
				fmt.Print("Step: ")
				nText := scanToken(reader)
				n, convErr := strconv.Atoi(nText)
				if convErr != nil {
					fmt.Println("Input step tidak valid")
					continue
				}
				if n >= 0 && n < len(pos) {
					step = n
				} else {
					fmt.Println("Step di luar rentang")
				}
			case "Q":
				quit = true
			default:
				fmt.Println("Input tidak valid")
			}
		}
		break
	}

	for {
		fmt.Print(">> Apakah Anda ingin menyimpan solusi? (Y/N) : ")
		sv := strings.ToUpper(scanToken(reader))
		if sv == "N" {
			break
		}
		if sv != "Y" {
			fmt.Println("Input tidak valid")
			continue
		}

		fmt.Print(">> Masukan nama file untuk menyimpan solusi : ")
		fn := scanToken(reader)
		if fn == "" {
			fn = "solusi.txt"
		}
		outputDir := filepath.Join("test", "output")
		if err := os.MkdirAll(outputDir, 0o755); err != nil {
			fmt.Println("Gagal membuat folder output:", err)
			break
		}
		outputFile := filepath.Join(outputDir, filepath.Base(fn))
		if errSv := saveSolution(outputFile, board, startPos, path, pos, cost, iter, dT.Milliseconds()); errSv != nil {
			fmt.Println("Gagal menyimpan:", errSv)
		} else {
			abs, _ := filepath.Abs(outputFile)
			fmt.Printf(">> Solusi disimpan pada %s\n", abs)
		}
		break
	}
}
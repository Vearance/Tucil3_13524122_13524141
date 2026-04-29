package main

import (
	"fmt"
	"time"
)

func main() {
	var filename string
	fmt.Print("Masukan file input : ")
	fmt.Scanln(&filename)

	var algo string
	fmt.Print("Algoritma apa yang anda pilih? (UCS/GBFS/A*)")
	fmt.Scanln(&algo)

	var heuristic string
	fmt.Print("Heuristic apa yang anda pilih? (H1/H2/H3)")
	fmt.Scanln(&heuristic)

	board, startPos, err := LoadMap(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	start := time.Now()
	// TODO: if for algo and heuristic
	// Solve()
	elapsed := time.Since(start)

	fmt.Printf("Waktu eksekusi: %d ms\n", elapsed.Milliseconds())
}
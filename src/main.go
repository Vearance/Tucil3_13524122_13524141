package main

import (
	"fmt"
	"time"
)

func main() {
	var filename string
	fmt.Print("Masukan file input : ")
	fmt.Scanln(&filename)

	board, startPos, err := LoadMap(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	start := time.Now()
	// Solve()
	elapsed := time.Since(start)

	fmt.Printf("Waktu eksekusi: %d ms\n", elapsed.Milliseconds()) [cite: 89, 106]
}
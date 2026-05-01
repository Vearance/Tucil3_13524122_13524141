package main

import (
	"fmt"
	"os"
	"strings"
)

func drawBoard(board *Board, startPos, at Point) string {
	var sb strings.Builder
	for i := 0; i < board.N; i++ {
		for j := 0; j < board.M; j++ {
			if i == at.X && j == at.Y {
				sb.WriteRune('Z')
			} else if i == startPos.X && j == startPos.Y {
				sb.WriteRune('*')
			} else {
				sb.WriteRune(board.Grid[i][j])
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func printSolution(board *Board, startPos Point, path string, pos []Point, cost int) {
	fmt.Printf("Solusi Yang Ditemukan : %s\n", path)
	fmt.Printf("Cost dari Solusi : %d\n\n", cost)
	fmt.Println("Initial")
	fmt.Print(drawBoard(board, startPos, pos[0]))
	for i := 1; i < len(pos); i++ {
		fmt.Println()
		fmt.Printf("Step %d : %c\n", i, path[i-1])
		fmt.Print(drawBoard(board, startPos, pos[i]))
	}
}

func saveSolution(fname string, board *Board, startPos Point, path string, pos []Point, cost, iter int, elapsed int64) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, "Solusi Yang Ditemukan : %s Cost dari Solusi : %d\n\n", path, cost)
	fmt.Fprintln(f, "Initial")
	fmt.Fprint(f, drawBoard(board, startPos, pos[0]))
	for i := 1; i < len(pos); i++ {
		fmt.Fprintln(f)
		fmt.Fprintf(f, "Step %d : %c\n", i, path[i-1])
		fmt.Fprint(f, drawBoard(board, startPos, pos[i]))
	}
	fmt.Fprintf(f, "\nWaktu eksekusi: %d ms\n", elapsed)
	fmt.Fprintf(f, "Banyak iterasi yang dilakukan: %d iterasi\n", iter)
	return nil
}

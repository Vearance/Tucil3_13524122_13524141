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

func printSolution(board *Board, startPos Point, r res) {
	fmt.Printf("Solusi Yang Ditemukan : %s Cost dari Solusi : %d\n\n", r.Path, r.Cost)
	fmt.Println("Initial")
	fmt.Print(drawBoard(board, startPos, r.Positions[0]))
	for i := 1; i < len(r.Positions); i++ {
		fmt.Println()
		fmt.Printf("Step %d : %c\n", i, r.Path[i-1])
		fmt.Print(drawBoard(board, startPos, r.Positions[i]))
	}
}

func saveSolution(fname string, board *Board, startPos Point, r res, elapsed int64) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, "Solusi Yang Ditemukan : %s Cost dari Solusi : %d\n\n", r.Path, r.Cost)
	fmt.Fprintln(f, "Initial")
	fmt.Fprint(f, drawBoard(board, startPos, r.Positions[0]))
	for i := 1; i < len(r.Positions); i++ {
		fmt.Fprintln(f)
		fmt.Fprintf(f, "Step %d : %c\n", i, r.Path[i-1])
		fmt.Fprint(f, drawBoard(board, startPos, r.Positions[i]))
	}
	fmt.Fprintf(f, "\nWaktu eksekusi: %d ms\n", elapsed)
	fmt.Fprintf(f, "Banyak iterasi yang dilakukan: %d iterasi\n", r.Iter)
	return nil
}

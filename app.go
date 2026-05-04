package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"Tucil3_13524122_13524141/src"
	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx       context.Context
	board     *solver.Board
	startPos  solver.Point
	path      string
	pos       []solver.Point
	cost      int
	iter      int
	elapsedMs int64
}

type LoadResult struct {
	OK    bool     `json:"ok"`
	Rows  int      `json:"rows"`
	Cols  int      `json:"cols"`
	Grid  []string `json:"grid"`
	Start [2]int   `json:"start"`
	Error string   `json:"error"`
}

type SolveResult struct {
	Found     bool     `json:"found"`
	Path      string   `json:"path"`
	Cost      int      `json:"cost"`
	Iter      int      `json:"iter"`
	ElapsedMs int64    `json:"elapsedMs"`
	Positions [][2]int `json:"positions"`
	Error     string   `json:"error"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) PickFile() string {
	path, _ := wailsruntime.OpenFileDialog(a.ctx, wailsruntime.OpenDialogOptions{
		Title: "Pilih file map",
	})
	return path
}

func (a *App) LoadFile(path string) LoadResult {
	if path == "" {
		return LoadResult{Error: "Path kosong"}
	}
	if _, err := os.Stat(path); err != nil {
		fb := filepath.Join("test", filepath.Base(path))
		if _, err2 := os.Stat(fb); err2 == nil {
			path = fb
		}
	}
	board, startPos, err := solver.LoadMap(path)
	if err != nil {
		return LoadResult{Error: err.Error()}
	}
	a.board = board
	a.startPos = startPos
	a.path = ""
	a.pos = nil

	rows := make([]string, board.N)
	for i := 0; i < board.N; i++ {
		rows[i] = string(board.Grid[i])
	}
	return LoadResult{
		OK:    true,
		Rows:  board.N,
		Cols:  board.M,
		Grid:  rows,
		Start: [2]int{startPos.X, startPos.Y},
	}
}

func (a *App) Solve(algo string, heuristic string) SolveResult {
	if a.board == nil {
		return SolveResult{Error: "Belum ada peta. Load file dulu."}
	}

	var hF func(solver.Point, *solver.Board, int) int
	switch heuristic {
	case "H1":
		hF = solver.ManhattanDistance
	case "H2":
		hF = solver.EuclideanDistance
	case "H3":
		hF = solver.ChebyshevDistance
	}

	t0 := time.Now()
	var goal *solver.State
	var iter int
	switch algo {
	case "UCS":
		goal, iter = solver.UCS(a.board, a.startPos)
	case "BFS":
		goal, iter = solver.BFS(a.board, a.startPos)
	case "GBFS":
		if hF == nil {
			return SolveResult{Error: "GBFS butuh heuristic"}
		}
		goal, iter = solver.GBFS(a.board, a.startPos, hF)
	case "A*":
		if hF == nil {
			return SolveResult{Error: "A* butuh heuristic"}
		}
		goal, iter = solver.AStar(a.board, a.startPos, hF)
	case "IDA*":
		if hF == nil {
			return SolveResult{Error: "IDA* butuh heuristic"}
		}
		goal, iter = solver.IDAStar(a.board, a.startPos, hF)
	default:
		return SolveResult{Error: "Algoritma tidak dikenal"}
	}
	a.elapsedMs = time.Since(t0).Milliseconds()
	a.iter = iter

	if goal == nil {
		a.pos = nil
		return SolveResult{
			Found:     false,
			Iter:      iter,
			ElapsedMs: a.elapsedMs,
		}
	}

	a.path, a.pos = solver.Reverse(goal)
	a.cost = goal.TotalCost

	positions := make([][2]int, len(a.pos))
	for i, p := range a.pos {
		positions[i] = [2]int{p.X, p.Y}
	}

	return SolveResult{
		Found:     true,
		Path:      a.path,
		Cost:      a.cost,
		Iter:      iter,
		ElapsedMs: a.elapsedMs,
		Positions: positions,
	}
}

func (a *App) Save(defaultName string) string {
	if a.pos == nil {
		return "Belum ada solusi"
	}
	if defaultName == "" {
		defaultName = "solusi.txt"
	}
	path, err := wailsruntime.SaveFileDialog(a.ctx, wailsruntime.SaveDialogOptions{
		Title:           "Simpan solusi",
		DefaultFilename: defaultName,
	})
	if err != nil {
		return fmt.Sprintf("Dialog error: %v", err)
	}
	if path == "" {
		return ""
	}
	if err := solver.SaveSolution(path, a.board, a.startPos, a.path, a.pos, a.cost, a.iter, a.elapsedMs); err != nil {
		return fmt.Sprintf("Gagal menyimpan: %v", err)
	}
	return "Saved to " + path
}

package main

type res struct {
	Found     bool
	Path      string
	Cost      int
	Iter      int
	Positions []Point
}

type key struct {
	X, Y, T int
}

func isGoal(s *State, b *Board) bool {
	if b.Grid[s.Pos.X][s.Pos.Y] != 'O' {
		return false
	}
	if b.MaxTarget == -1 {
		return true
	}
	return s.CurrentTarget > b.MaxTarget
}

func reconstruct(goal *State) (string, []Point) {
	n := 0
	for s := goal; s != nil; s = s.Parent {
		n++
	}
	ps := make([]Point, n)
	dirs := make([]byte, n-1)
	i := n - 1
	for s := goal; s != nil; s = s.Parent {
		ps[i] = s.Pos
		if s.Parent != nil {
			dirs[i-1] = s.Dir
		}
		i--
	}
	return string(dirs), ps
}

func UCS(board *Board, startPos Point) res {
	pq := &PriorityQueue{}

	visited := make(map[key]int)
	dirs := []string{"U", "D", "L", "R"}

	first := &State{Pos: startPos, TotalCost: 0, CurrentTarget: 0}
	pq.Push(&Item{state: first, priority: 0})

	iter := 0
	for pq.Len() > 0 {
		it := pq.Pop()
		cur := it.state
		iter++

		if isGoal(cur, board) {
			path, ps := reconstruct(cur)
			return res{true, path, cur.TotalCost, iter, ps}
		}

		k := key{cur.Pos.X, cur.Pos.Y, cur.CurrentTarget}
		prev, ok := visited[k]
		if ok && prev <= cur.TotalCost {
			continue
		}
		visited[k] = cur.TotalCost

		for i := 0; i < len(dirs); i++ {
			next, ok := Move(board, cur, dirs[i])
			if !ok {
				continue
			}
			nk := key{next.Pos.X, next.Pos.Y, next.CurrentTarget}
			pv, vok := visited[nk]
			if vok && pv <= next.TotalCost {
				continue
			}
			pq.Push(&Item{state: next, priority: next.TotalCost})
		}
	}

	return res{false, "", 0, iter, nil}
}

func GBFS(board *Board, startPos Point, h func(Point, *Board, int) int) res {
	pq := &PriorityQueue{}

	visited := make(map[key]bool)
	dirs := []string{"U", "D", "L", "R"}

	first := &State{Pos: startPos, TotalCost: 0, CurrentTarget: 0}
	pq.Push(&Item{state: first, priority: h(startPos, board, 0)})

	iter := 0
	for pq.Len() > 0 {
		it := pq.Pop()
		cur := it.state
		iter++

		if isGoal(cur, board) {
			path, ps := reconstruct(cur)
			return res{true, path, cur.TotalCost, iter, ps}
		}

		k := key{cur.Pos.X, cur.Pos.Y, cur.CurrentTarget}
		if visited[k] {
			continue
		}
		visited[k] = true

		for i := 0; i < len(dirs); i++ {
			next, ok := Move(board, cur, dirs[i])
			if !ok {
				continue
			}
			if visited[key{next.Pos.X, next.Pos.Y, next.CurrentTarget}] {
				continue
			}
			hv := h(next.Pos, board, next.CurrentTarget)
			pq.Push(&Item{state: next, priority: hv})
		}
	}

	return res{false, "", 0, iter, nil}
}

func AStar(board *Board, startPos Point, h func(Point, *Board, int) int) res {
	pq := &PriorityQueue{}

	visited := make(map[key]int)
	dirs := []string{"U", "D", "L", "R"}

	first := &State{Pos: startPos, TotalCost: 0, CurrentTarget: 0}
	pq.Push(&Item{state: first, priority: h(startPos, board, 0)})

	iter := 0
	for pq.Len() > 0 {
		it := pq.Pop()
		cur := it.state
		iter++

		if isGoal(cur, board) {
			path, ps := reconstruct(cur)
			return res{true, path, cur.TotalCost, iter, ps}
		}

		k := key{cur.Pos.X, cur.Pos.Y, cur.CurrentTarget}
		prev, ok := visited[k]
		if ok && prev <= cur.TotalCost {
			continue
		}
		visited[k] = cur.TotalCost

		for i := 0; i < len(dirs); i++ {
			next, ok := Move(board, cur, dirs[i])
			if !ok {
				continue
			}
			nk := key{next.Pos.X, next.Pos.Y, next.CurrentTarget}
			pv, vok := visited[nk]
			if vok && pv <= next.TotalCost {
				continue
			}
			hv := h(next.Pos, board, next.CurrentTarget)
			f := next.TotalCost + hv
			pq.Push(&Item{state: next, priority: f})
		}
	}

	return res{false, "", 0, iter, nil}
}
